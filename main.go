package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/ping", ping)

	r.Run()
}

func ping(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
