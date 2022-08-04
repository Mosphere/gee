package gee

import (
	"fmt"
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		//Start timer
		t := time.Now()
		//Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func V1() HandlerFunc {
	return func(ctx *Context) {
		//t := time.Now()
		fmt.Println("v1 before")

		ctx.Next()

		fmt.Println("v1 after")
		//log.Printf("[%d] %s in %v", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}
