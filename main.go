package main

import "github.com/gin-gonic/gin"

const port = ":9090"

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	r.Run()
}
