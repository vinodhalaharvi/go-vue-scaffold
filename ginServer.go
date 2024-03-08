package main

import (
	"github.com/gin-gonic/gin"
)

type GinServer struct {
	Router *gin.Engine
}

func NewGinServer(router *gin.Engine) *GinServer {
	return &GinServer{Router: router}
}

func (g *GinServer) Static() {
	g.Router.Static("/web", "./web/dist")
}

func (g *GinServer) LoadHTMLGlob() {
	g.Router.LoadHTMLGlob("templates/*")
}

func (g *GinServer) Setup() {
	g.Static()
	g.LoadHTMLGlob()
	g.InstallRoutes()
}

func (g *GinServer) InstallRoutes() {
	g.Router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

}

func (g *GinServer) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (g *GinServer) Run() error {
	return g.Router.Run(":8080")
}
