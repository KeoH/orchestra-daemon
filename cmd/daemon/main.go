package main

import (
	"github.com/savsgio/atreugo/v7"
)

func main() {

	config := &atreugo.Config{
		Host: "0.0.0.0",
		Port: 12001,
	}

	server := atreugo.New(config)

	server.Path("GET", "/", func(ctx *atreugo.RequestCtx) error {
		return ctx.JSONResponse(atreugo.JSON{"Daemon": "Ok"})
	})

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
