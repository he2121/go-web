package hehe

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	split := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, part := range split {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}

	return parts
}

func (router *router) addRouter(method, path string, handler HandlerFunc) {
	parts := parsePattern(path)
	_, ok := router.roots[method]
	if !ok {
		router.roots[method] = new(node)
	}
	router.roots[method].insert(path, parts, 0)
	router.handlers[method+"-"+path] = handler
}

func (router *router) getRouter(method, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)

	root, ok := router.roots[method]
	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)
	if n == nil {
		return nil, nil
	}
	parts := parsePattern(n.pattern)

	params := make(map[string]string, 0)
	for i, part := range parts {
		if part[0] == ':' {
			params[part[1:]] = searchParts[i]
			continue
		}
		if part[0] == '*' && len(part) > 1 {
			params[part[1:]] = strings.Join(searchParts[i:], "/")
			break
		}
	}
	return n, params
}

func (router *router) handle(c *Context) {
	n, params := router.getRouter(c.Method, c.Path)
	if n != nil {
		c.Params = params
		router.handlers[c.Method+"-"+n.pattern](c)
	} else {
		c.HTML(http.StatusNotFound, "404 NOT FOUND: "+c.Path)
	}
}
