package gee

import (
	"net/http"
)

type HandlerFunc func(*Context)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	router      *Router
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	return &RouterGroup{
		prefix: group.prefix + prefix,
		router: group.router,
	}
}

func (group *RouterGroup) addRoute(method, pattern string, handler HandlerFunc) {
	pattern = group.prefix + pattern
	group.router.addRoute(method, pattern, handler)
}

//Get defines the method to add get route
func (group *RouterGroup) Get(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

//Post defines the method to add post route
func (group *RouterGroup) Post(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

type Engine struct {
	*RouterGroup
}

func New() *Engine {
	engine := &Engine{}
	engine.RouterGroup = &RouterGroup{router: newRouter()}
	return engine
}

//Run defines the method to start a http server
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

//ServerHTTP implents the method ServerHTTP in http.Handler interface
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	engine.router.handle(c)
}
