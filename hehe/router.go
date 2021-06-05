package hehe

import (
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router{
	return &router{handlers: map[string]HandlerFunc{}}
}

func (router *router) addRouter(method, path string, handler HandlerFunc) {
	router.handlers[method+"-"+path] = handler
}

func (router *router) handle(c *Context)  {
	key := c.Method + "-" + c.Path
	if handler, ok := router.handlers[key]; ok {
		handler(c)
	} else {
		c.HTML(http.StatusNotFound, "404 NOT FOUND: " + c.Path)
	}
}

