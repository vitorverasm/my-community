package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	Initialize()
}

func Initialize() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000")
	log.Println("Running server on port 3000...")
}
