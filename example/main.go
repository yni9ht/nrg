package main

import (
	"github.com/yni9ht/nrg"
	"log"
	"net/http"
)

func main() {
	server := nrg.NewServer()

	server.GET("/ping", func(context *nrg.Context) {
		if id, ok := context.GetQuery("id"); ok {
			context.JSON(http.StatusOK, id)
			return
		}
		context.JSON(http.StatusOK, "pong")
	})

	server.POST("/post", func(context *nrg.Context) {
		context.JSON(http.StatusOK, "post")
	})

	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}
