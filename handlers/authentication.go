package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gotrueTypes "github.com/supabase-community/gotrue-go/types"
	supabaseSdk "github.com/supabase-community/supabase-go"
	"github.com/vitorverasm/my-community/pkg/stream"
	"github.com/vitorverasm/my-community/types"
)

func HandleLogin(c *gin.Context, sp *supabaseSdk.Client) {
	var loginRequestBody types.LoginRequestBody
	c.BindJSON(&loginRequestBody)
	token, signInError := sp.Auth.SignInWithEmailPassword(loginRequestBody.Email, loginRequestBody.Password)

	authorizedClient := sp.Auth.WithToken(
		token.AccessToken,
	)

	user, getUserError := authorizedClient.GetUser()

	if getUserError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": getUserError.Error(), "msg": "Failed to get user information"})
		return
	}

	if signInError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": signInError.Error(), "msg": "Login failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": loginRequestBody.Email, "token": token.AccessToken, "streamToken": user.UserMetadata["streamToken"].(string)})
}

func HandleSignUp(c *gin.Context, sp *supabaseSdk.Client) {
	var signUpRequestBody types.SignUpRequestBody
	c.BindJSON(&signUpRequestBody)
	token, err := stream.GetToken(signUpRequestBody.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "msg": "Failed to create token for user"})
	}

	res, error := sp.Auth.Signup(gotrueTypes.SignupRequest{
		Email:    signUpRequestBody.Email,
		Password: signUpRequestBody.Password,
		Data: map[string]interface{}{
			"streamToken": token,
		},
	})
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error(), "msg": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": res.User.Email, "streamToken": token, "msg": "Email confirmation was sent to your email address"})
}
