package launch

import (
	"encoding/json"
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"
)

const (
	ApplicationJSON   = "application/json"
	MultipartFormData = "multipart/form-data"
)

var plugins []any

type Validator interface {
	Validate() error
}

type Handler func(*Ctx) error

type Plugin func(Handler) Handler

type ErrorHandler func(*Ctx, error)

type Ctx struct {
	r            *http.Request
	w            http.ResponseWriter
	params       httprouter.Params
	status       int
	App          *Launch
	ErrorHandler ErrorHandler
	RequestID    string
}

func (c *Ctx) Request() *http.Request {
	return c.r
}

func (c *Ctx) Param(name string) string {
	return c.params.ByName(name)
}

func (c *Ctx) Status(s int) *Ctx {
	c.status = s
	return c
}

func (c *Ctx) JSON(v any) error {
	c.w.WriteHeader(c.status)
	c.w.Header().Add("Content-Type", ApplicationJSON)
	return json.NewEncoder(c.w).Encode(v)
}

type Launch struct {
	routePrefix  string
	router       *httprouter.Router
	plugins      []Plugin
	errorHandler ErrorHandler
}

var App = &Launch{
	router:       httprouter.New(),
	plugins:      []Plugin{},
	errorHandler: func(c *Ctx, err error) {},
}

func New() *Launch {
	return &Launch{
		router:  httprouter.New(),
		plugins: []Plugin{},
	}
}

func (l *Launch) makeHTTPHandler(h Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := &Ctx{
			r:            r,
			w:            w,
			params:       p,
			status:       http.StatusOK,
			ErrorHandler: defaultErrorHandler,
		}
		for i := len(l.plugins) - 1; i >= 0; i-- {
			h = l.plugins[i](h)
		}
		if err := h(ctx); err != nil {
			ctx.ErrorHandler(ctx, err)
			return
		}
	}
}

func (l *Launch) addHandler(method, route string, h Handler) {
	route = path.Join(l.routePrefix, route)
	l.router.Handle(method, route, l.makeHTTPHandler(h))
}

func (l *Launch) Plug(plugins ...Plugin) {
	l.plugins = append(l.plugins, plugins...)
}

func (l *Launch) Get(route string, h Handler) {
	l.addHandler("GET", route, h)
}

func (l *Launch) Post(route string, h Handler) {
	l.addHandler("POST", route, h)
}

func (l *Launch) Put(route string, h Handler)     {}
func (l *Launch) Patch(route string, h Handler)   {}
func (l *Launch) Delete(route string, h Handler)  {}
func (l *Launch) Head(route string, h Handler)    {}
func (l *Launch) Options(route string, h Handler) {}

func (l *Launch) Start() {
	http.ListenAndServe(":3000", l.router)
}

func RequestParams[T any](c *Ctx) (T, error) {
	var params T
	if err := json.NewDecoder(c.r.Body).Decode(&params); err != nil {
		return params, err
	}
	if err := validateRequestParams(params); err != nil {
		return params, err
	}
	return params, nil
}

func validateRequestParams(params any) error {
	if v, ok := params.(Validator); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func defaultErrorHandler(c *Ctx, err error) {
	c.Status(http.StatusInternalServerError).JSON(map[string]string{
		"error": err.Error(),
	})
}
