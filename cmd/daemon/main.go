package main

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/savsgio/atreugo/v7"
)

func main() {

	config := &atreugo.Config{
		Host: "0.0.0.0",
		Port: 12001,
	}

	server := atreugo.New(config)

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	server.Path("GET", "/", func(ctx *atreugo.RequestCtx) error {
		return ctx.JSONResponse(atreugo.JSON{"status": "Ok"})
	})

	server.Path("GET", "/containers", func(ctx *atreugo.RequestCtx) error {
		containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
		if err != nil {
			panic(err)
		}
		return ctx.JSONResponse(atreugo.JSON{"containers": containers, "status": "Ok"})
	})

	servError := server.ListenAndServe()
	if servError != nil {
		panic(servError)
	}
}
