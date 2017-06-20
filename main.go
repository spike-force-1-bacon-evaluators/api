package main

import (
	"github.com/spike-force-1-bacon-evaluators/api/handlers"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	r := gin.Default()
	// r.LoadHTMLGlob("templates/*")
	r.GET("/", handlers.Index())
	r.GET("/map", handlers.Map())
	r.Run()
}
