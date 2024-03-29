package launch

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"sync"

	"github.com/julienschmidt/httprouter"
)

const (
	ApplicationJSON   = "application/json"
	MultipartFormData = "multipart/form-data"
)

func isMultipartFormData(contentType string) bool {
	return true
}

func parseFormData(c *Ctx, params any) error {
	return nil
}

var plugins []any

var ctxPool = sync.Pool{
	New: newCtx,
}

type Validator interface {
	Validate() error
}

type Handler func(*Ctx) error

type Plugin func(Handler) Handler

func applyPlugins(h Handler, plugins ...Plugin) Handler {
	for i := len(plugins) - 1; i >= 0; i-- {
		h = plugins[i](h)
	}
	return h
}

type ErrorHandler func(*Ctx, error)

type Ctx struct {
	RequestID    string
	Context      context.Context
	Request      *http.Request
	Response     http.ResponseWriter
	ErrorHandler ErrorHandler

	params httprouter.Params
	status int
	launch *Launch
}

func newCtx() any {
	return &Ctx{
		status:  http.StatusOK,
		Context: context.Background(),
	}
}

func (c *Ctx) SetNotFound(h Handler) error {
	c.launch.router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := h(c); err != nil {
			c.ErrorHandler(c, err)
		}
	})
	return nil
}

func (c *Ctx) reset() {
	c.Request = nil
	c.Response = nil
	c.launch = nil
}

func (c *Ctx) init(w http.ResponseWriter, r *http.Request, p httprouter.Params, l *Launch) {
	c.Response = w
	c.Request = r
	c.params = p
	c.launch = l
	c.Context = context.Background()
}

func (c *Ctx) Header(name string) string {
	return c.Request.Header.Get(name)
}

func (c *Ctx) SetHeader(key, value string) {
	c.Response.Header().Set(key, value)
}

func (c *Ctx) Form(name string) string {
	return c.Request.FormValue(name)
}

func (c *Ctx) Query(name string) string {
	return c.Request.URL.Query().Get(name)
}

func (c *Ctx) Param(name string) string {
	return c.params.ByName(name)
}

func (c *Ctx) Status(s int) *Ctx {
	c.status = s
	return c
}

func (c *Ctx) JSON(v any) error {
	c.Response.Header().Add("Content-Type", ApplicationJSON)
	c.Response.WriteHeader(c.status)
	return json.NewEncoder(c.Response).Encode(v)
}

func (c *Ctx) View(name string, data any) error {
	return c.launch.renderer.Render(c.Response, name, data)
}

type Router struct {
	Launch
}

type Launch struct {
	routePrefix string
	router      *httprouter.Router
	plugins     []Plugin
	renderer    ViewRenderer
}

var App = &Launch{
	router:  httprouter.New(),
	plugins: []Plugin{},
}

func New(onlyAPI bool) *Launch {
	var renderer ViewRenderer
	if !onlyAPI {
		// panic("This is triggered")
		renderer = *NewViewRenderer("www" + "/" + "views")
	}
	return &Launch{
		router:   httprouter.New(),
		plugins:  []Plugin{},
		renderer: renderer,
	}
}

func (l *Launch) Router(prefix string, plugins ...Plugin) *Router {
	router := &Router{*l}
	router.Launch.routePrefix = prefix
	return router
}

func (l *Launch) makeHTTPHandler(h Handler) httprouter.Handle {
	h = errorHandlerPlug(h)
	h = requestIdPlug(h)
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := ctxPool.Get().(*Ctx)
		ctx.Request = r
		ctx.Response = w
		ctx.launch = l
		ctx.params = p
		ctx.Context = context.Background()
		handler := applyPlugins(h, l.plugins...)
		if err := handler(ctx); err != nil {
			ctx.ErrorHandler(ctx, err)
		}
		ctx.reset()
		ctxPool.Put(ctx)
	}
}

func (l *Launch) addHandler(method, route string, h Handler) {
	route = path.Join(l.routePrefix, route)
	handler := l.makeHTTPHandler(h)
	l.router.Handle(method, route, handler)
}

func (l *Launch) Plug(plugins ...Plugin) {
	l.plugins = append(l.plugins, plugins...)
}

func (l *Launch) Get(route string, h Handler) {
	fmt.Println("test all the things")
	l.addHandler("GET", route, h)
}

func (l *Launch) Post(route string, h Handler) {
	l.addHandler("POST", route, h)
}

func (l *Launch) Put(route string, h Handler) {
	l.addHandler("PUT", route, h)
}
func (l *Launch) Patch(route string, h Handler) {
	l.addHandler("PATCH", route, h)
}
func (l *Launch) Delete(route string, h Handler) {
	l.addHandler("DELETE", route, h)
}
func (l *Launch) Head(route string, h Handler) {
	l.addHandler("HEAD", route, h)
}
func (l *Launch) Options(route string, h Handler) {
	l.addHandler("OPTIONS", route, h)
}

func (l *Launch) Start(listenAddr string) {
	fmt.Printf("app running on 127.0.0.1:%s\n", listenAddr)
	http.ListenAndServe(listenAddr, l.router)
}

func RequestParams[T any](c *Ctx) (T, error) {
	var params T
	contentType := c.Header("Content-Type")
	if contentType == "application/json" {
		if err := json.NewDecoder(c.Request.Body).Decode(&params); err != nil {
			return params, err
		}
	}
	if isMultipartFormData(contentType) {
		if err := parseFormData(c, &params); err != nil {
			return params, err
		}
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
