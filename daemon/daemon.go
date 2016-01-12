package main

import (
	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func main() {
	r := gin.Default()
	r.GET("/", indexHandler)

	r.Run(":4567")
}
