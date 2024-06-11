package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vitorverasm/my-community/handlers"
	"github.com/vitorverasm/my-community/pkg/supabase"
)

var sp = supabase.InitializeClient()

func main() {
	InitializeAPI()
}

func InitializeAPI() {
	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		handlers.HandleLogin(c, sp)
	})

	r.POST("/register", func(c *gin.Context) {
		handlers.HandleSignUp(c, sp)
	})

	r.Run(":3000")
	log.Println("Running server on port 3000...")
}
