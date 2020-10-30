package main

import "github.com/gin-gonic/gin"

func httpGET(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()
	r.GET("/", httpGET)
	r.Run()
}
