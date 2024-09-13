package main

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
)

var pool = sync.Pool{
	New: newCtx,
}

type Validator interface {
	Validate() error
}

type Handler func(ctx *Ctx) error

type Plugin func(h Handler) Handler

type ErrorHandler func(ctx *Ctx, err error)

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

func (c *Ctx) reset() {
  c.Request = nil
  c.Response = nil
  c.launch = nil
}

func _newCtx(l *Launch, w http.ResponseWriter, r *http.Request, p httprouter.Params) *Ctx {
	return &Ctx{
		Request:  r,
		Response: w,
		params:   p,
		launch:   l,
	}
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

func (c *Ctx) Status(status int) *Ctx {
	c.status = status
	return c
}

func (c *Ctx) JSON(v any) error {
	c.Response.Header().Set("Content-Type", "application/json")
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
	router       *httprouter.Router
	routerPrefix string
	plugins      []Plugin
	renderer     Renderer
}

type Renderer struct {
}

func NewViewRenderer(pathname string) Renderer {
	return Renderer{}
}

func (r *Renderer) Render(w http.ResponseWriter, name string, data any) error {
	return nil
}

func New(onlyAPI bool) *Launch {
	var renderer Renderer
	if !onlyAPI {
		renderer = NewViewRenderer("www" + "/" + "views")
	}
	return &Launch{
		router:   httprouter.New(),
		plugins:  []Plugin{},
		renderer: renderer,
	}
}

func (l *Launch) Router(prefix string, plugins ...Plugin) *Router {
	router := &Router{*l}
	router.Launch.routerPrefix = prefix
	return router
}

func (l *Launch) makeHTTPHandler(h Handler) httprouter.Handle {
	h = errorHandlePlug(h)
	h = requestIDPlug(h)
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// ctx := newCtx(l, w, r, p)
		ctx := pool.Get().(*Ctx)
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
		pool.Put(ctx)
	}
}

func applyPlugins(h Handler, plugins ...Plugin) Handler {
	for i := len(plugins) - 1; i >= 0; i-- {
		h = plugins[i](h)
	}
	return h
}

func (l *Launch) addHandler(method, route string, h Handler) {

}

func (l *Launch) Get(pathname string, lh Handler) {

}

func errorHandlePlug(h Handler) Handler {
	return h
}

func requestIDPlug(h Handler) Handler {
	return h
}
