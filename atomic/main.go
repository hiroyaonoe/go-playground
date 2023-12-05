package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

type Context struct {
	Request *http.Request
}

func main() {
	fmt.Println("Hello, World!")
	c := &Context{
		Request: &http.Request{},
	}

	go func() {
		_ = c.Request != nil && c.Request.Context() != nil
	}()

	go func(c *Context) {
		// req := c.Request.WithContext(context.Background())
		// c.Request = req
		c.Request = c.Request.WithContext(context.Background())
		// c.Request.WithContext(context.Background())
	}(c)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "1")

	req := &http.Request{}
	req = req.WithContext(ctx)
	fmt.Println(req.Context().Value("key"))

	ctx = context.Background()
	ctx = context.WithValue(ctx, "key", "2")
	req = req.WithContext(ctx)
	fmt.Println(req.Context().Value("key"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
