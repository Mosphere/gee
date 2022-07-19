package main

import (
	"fmt"
	"geeweb/gee"
	"net/http"
	"strings"
)

func main() {
	r := gee.New()
	v1 := r.Group("v1")
	{
		v1.Get("/", func(c *gee.Context) {
			c.Html(http.StatusOK, "<p>hello</p>")
		})

		v1.Get("/hello", func(c *gee.Context) {
			parsePattern(c.Path)
			c.String(http.StatusOK, "hello %s, you are at %s\n", c.Query("name"), c.Path)

		})
	}

	r.Run(":80")
}

func parsePattern(path string) {
	arr := strings.Split(path, "/")
	fmt.Println(arr)
}
