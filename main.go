package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.Get("/", indexHandler)

	r.Run(":80")
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}
