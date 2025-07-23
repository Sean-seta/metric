package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"
	"time"
)

func main() {
	r := gin.New()

	// NewWithConfig is the recommended way to initialize the middleware
	p := ginprometheus.NewWithConfig(ginprometheus.Config{
		Subsystem: "dat",
	})

	p.Use(r)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello world!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/health", func(c *gin.Context) {
		time.Sleep(2 * time.Second)
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	r.Run(":8080")
}
