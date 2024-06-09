package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	supabaseSdk "github.com/supabase-community/supabase-go"
	"github.com/vitorverasm/my-community/types"
)

func HandleLoginWithSupabase(c *gin.Context, sp *supabaseSdk.Client) {
	var loginRequestBody types.LoginRequestBody
	c.BindJSON(&loginRequestBody)

	token, error := sp.Auth.SignInWithEmailPassword(loginRequestBody.Email, loginRequestBody.Password)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": loginRequestBody.Email, "token": token.AccessToken})
}
