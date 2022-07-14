package main

import (
	"fmt"
	"gee"
	"net/http"
	"strings"
)

func main() {
	r := gee.New()

	//r.Get("/", indexHandler)
	r.Get("/hello", func(c *gee.Context) {
		parsePattern(c.Path)
		c.String(http.StatusOK, "hello %s, you are at %s\n", c.Query("name"), c.Path)
		//c.Html(http.StatusOK, "<p>hello</p>")
	})

	r.Run(":80")
}

func parsePattern(path string) {
	arr := strings.Split(path, "/")
	fmt.Println(arr)
}
