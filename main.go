package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"
	"math/rand"
	"time"
)

func genRandom() time.Duration {
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 5
	randomNumber := rand.Intn(300) + 1 // 1 to 300

	return time.Duration(randomNumber) * time.Millisecond
}

func main() {
	r := gin.New()

	// NewWithConfig is the recommended way to initialize the middleware
	p := ginprometheus.NewWithConfig(ginprometheus.Config{
		Subsystem: "gin",
	})

	p.Use(r)
	r.GET("/", func(c *gin.Context) {
		start := time.Now()
		time.Sleep(genRandom())
		c.JSON(200, "Hello world!")
		end := time.Now()
		fmt.Printf("Path: /, Request duration: %f seconds\n", end.Sub(start).Seconds())
	})
	r.GET("/ping", func(c *gin.Context) {
		start := time.Now()
		time.Sleep(genRandom())
		c.JSON(200, gin.H{
			"message": "pong",
		})
		end := time.Now()
		fmt.Printf("Path: /ping, Request duration: %f seconds\n", end.Sub(start).Seconds())
	})
	r.GET("/health", func(c *gin.Context) {
		start := time.Now()
		time.Sleep(genRandom())
		c.JSON(200, gin.H{
			"status": "ok",
		})
		end := time.Now()
		fmt.Printf("Path: /health, Request duration: %f seconds\n", end.Sub(start).Seconds())
	})
	r.Run(":8080")
}
