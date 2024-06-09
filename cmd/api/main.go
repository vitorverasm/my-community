package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorverasm/my-community/pkg/supabase"
	"github.com/vitorverasm/my-community/types"
)

func main() {
	InitializeAPI()
}

func InitializeAPI() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login", func(c *gin.Context) {
		var loginRequestBody types.LoginRequestBody
		c.BindJSON(&loginRequestBody)

		supabaseClient := supabase.InitializeClient()

		token, error := supabaseClient.Auth.SignInWithEmailPassword(loginRequestBody.Email, loginRequestBody.Password)

		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"email": loginRequestBody.Email, "token": token.AccessToken})
	})

	r.Run(":3000")
	log.Println("Running server on port 3000...")
}
