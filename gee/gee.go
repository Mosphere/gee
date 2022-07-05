package gee

import (
	"fmt"
	"net/http"
)

type Engine struct {
	router map[string]http.HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]http.HandlerFunc)}
}

func (engine *Engine) addRoute(method, pattern string, handler http.HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

//Get defines the method to add get route
func (engine *Engine) Get(pattern string, handler http.HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

//Post defines the method to add post route
func (engine *Engine) Post(pattern string, handler http.HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

//Run defines the method to start a http server
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

//ServerHTTP implents the method ServerHTTP in http.Handler interface
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
