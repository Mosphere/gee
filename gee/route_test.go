package gee

import (
	"fmt"
	"testing"
)

func InitRouter() *Router {
	r := newRouter()
	r.addRoute("GET", "/p/:lang/doc", nil)
	r.addRoute("GET", "/asserts/*filepath", nil)
	return r
}

func TestGetRoute(t *testing.T) {
	r := InitRouter()
	n, params := r.getRoute("GET", "/p/go/doc")
	if n == nil {
		t.Fatal("nil should not be returned.")
	}

	if n.Pattern != "/p/:lang/doc" {
		t.Fatal("should match /p/:lang/doc.")
	}

	if params["lang"] != "go" {
		t.Fatal("lang should equal to 'go'.")
	}

	fmt.Printf("matched path: %s, params['lang']: %s \n", n.Pattern, params["lang"])
}
