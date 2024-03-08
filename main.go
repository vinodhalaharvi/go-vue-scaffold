package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	server := NewGinServer(router)
	server.Setup()
	err := server.Run()
	if err != nil {
		log.Fatalf("server.Run() failed with %s\n", err)
	}
}
