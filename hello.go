package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a router object
	r := gin.Default()

	// Return a status 200 code and a string hello world
	// for the "/hello" path
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// Create a group of routes
	api := r.Group("/api")

	// Return a JSON response
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Add static view support (some directory are required)
	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Listen on port 3000
	r.Run(":8181")
}
