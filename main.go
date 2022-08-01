package main

import (
	"fmt"
	"geeweb/gee"
	"net/http"
	"strings"
)

func main() {
	r := gee.Default()
	r.Get("/panic", func(c *gee.Context) {
		names := []string{"geetutu"}
		// panic because of index out of names
		c.String(http.StatusOK, names[50])
	})
	v1 := r.Group("/v1")
	v1.Use(gee.V1())
	{
		v1.Get("/hello", func(c *gee.Context) {
			//fmt.Println("test...")
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
