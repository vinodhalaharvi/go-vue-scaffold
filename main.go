package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	server := NewGinServer(router)
	server.Setup()
	fmt.Printf("http://localhost:8082/web/index.html\n")
	err := server.Run(":8082")
	if err != nil {
		log.Fatalf("server.Run() failed with %s\n", err)
	}
}
