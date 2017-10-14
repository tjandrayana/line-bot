package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tjandrayana/line-bot/messages"
)

func main() {
	r := gin.New()
	r.GET("/ping", ping)
	r.POST("/line/triger", messages.Triger)
	r.Run()
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
