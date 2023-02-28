package main

import (
	"github.com/yni9ht/nrg"
	"log"
	"net/http"
)

func main() {
	server := nrg.NewServer()

	server.Get("/ping", func(context *nrg.Context) {
		context.JSON(http.StatusOK, "pong")
	})

	server.POST("/post", func(context *nrg.Context) {
		context.JSON(http.StatusOK, "post")
	})

	if err := server.Run(":8181"); err != nil {
		log.Fatalln(err)
	}
}
