package launch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Validator interface {
	Validate() (any, bool)
}

type ErrorHandler func(*Ctx, error)

type PostHandler[T any] func(*PostCtx[T]) error

type Handler func(*Ctx) error

const (
	ApplicationJSON   = "application/json"
	MultipartFormData = "multipart/form-data"
)

func validateRequestParams(ctx *Ctx, v any) bool {
	if v, ok := v.(Validator); ok {
		if errs, ok := v.Validate(); !ok {
			ctx.Status(http.StatusBadRequest).JSON(errs)
			return false
		}
	}
	return true
}

type Ctx struct {
	r          *http.Request
	w          http.ResponseWriter
	statusCode int
	params     httprouter.Params
}

func newCtx(w http.ResponseWriter, r *http.Request, params httprouter.Params) *Ctx {
	return &Ctx{
		r:          r,
		w:          w,
		statusCode: http.StatusOK,
		params:     params,
	}
}

type Param string

func (p Param) AsInt(ctx *Ctx) (int, error) {
	val, err := strconv.Atoi(string(p))
	if err != nil {
		m := map[string]string{"error": fmt.Sprintf("param <%s> not of type <int>", p)}
		ctx.Status(http.StatusBadRequest).JSON(m)
		return 0, err
	}
	return val, nil
}

func (c *Ctx) Param(name string) Param {
	return Param(c.params.ByName(name))
}

func (c *Ctx) Status(s int) *Ctx {
	c.statusCode = s
	return c
}

func (c *Ctx) JSON(v any) error {
	c.w.WriteHeader(c.statusCode)
	c.w.Header().Add("Content-Type", ApplicationJSON)
	return json.NewEncoder(c.w).Encode(v)
}

type launch struct {
	errorHandler ErrorHandler
	router       *httprouter.Router
}

var App = &launch{
	errorHandler: func(ctx *Ctx, err error) {},
	router:       httprouter.New(),
}

type PostCtx[T any] struct {
	Ctx
	params T
}

func (c *PostCtx[T]) RequestParams() T {
	return c.params
}

func Start() {
	http.ListenAndServe(":3000", App.router)
}

func Get(path string, h Handler) {
	fn := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := newCtx(w, r, p)
		if r.Method != "GET" {
			ctx.Status(http.StatusMethodNotAllowed).JSON(map[string]string{
				"error": fmt.Sprintf("method <%s> not allowed", r.Method),
			})
			return
		}
		h(ctx)
	}
	App.router.GET(path, fn)
}

func Post[T any](path string, h PostHandler[T]) {
	fn := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var (
			params T
			ctx    = newCtx(w, r, p)
		)
		if r.Method != "POST" {
			ctx.Status(http.StatusMethodNotAllowed).JSON(map[string]string{
				"error": fmt.Sprintf("method <%s> not allowed", r.Method),
			})
			return
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			App.errorHandler(ctx, err)
			return
		}
		if !validateRequestParams(ctx, params) {
			return
		}
		h(&PostCtx[T]{
			Ctx:    *ctx,
			params: params,
		})
	}
	App.router.POST(path, fn)

}
