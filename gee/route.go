package gee

import (
	"net/http"
	"strings"
)

type Router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{handlers: make(map[string]HandlerFunc), roots: make(map[string]*node)}
}

//解析pattern，不保留开头的空字符串，如 /p/lang/c, ["p", "lang", "c"]
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, part := range vs {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}
func (r *Router) addRoute(method, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	node := r.roots[method]
	node.Insert(pattern, parts, 0)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// Get params from path
func (r *Router) getRoute(method string, path string) (*node, map[string]string) {
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	searchParts := parsePattern(path)
	n := root.Search(searchParts, 0)
	params := make(map[string]string, 0)
	if n != nil {
		parts := parsePattern(n.Pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}

			if part[0] == '*' && len(parts) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *Router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + c.Path
		handler := r.handlers[key]
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
