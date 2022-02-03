package main

import "github.com/gin-gonic/gin"

func main() {
	// Create a router object
	r := gin.Default()

	// Return a status 200 code and a string hello world
	// for the "/hello" path
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// Listen on port 3000
	r.Run(":3000")
}
