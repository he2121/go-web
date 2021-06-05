package hehe

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}


func (engine *Engine) GET(path string, handler HandlerFunc) {
	engine.router.addRouter(http.MethodGet, path, handler)
}

func (engine *Engine) POST(path string, handler HandlerFunc) {
	engine.router.addRouter(http.MethodPost, path, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
