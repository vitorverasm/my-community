package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gotrueTypes "github.com/supabase-community/gotrue-go/types"
	supabaseSdk "github.com/supabase-community/supabase-go"
	"github.com/vitorverasm/my-community/pkg/stream"
	"github.com/vitorverasm/my-community/types"
)

func NewLoginHandler(ap types.AuthProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		HandleLogin(c, ap)
	}
}

func HandleLogin(c *gin.Context, ap types.AuthProvider) {
	var loginRequestBody types.LoginRequestBody
	c.BindJSON(&loginRequestBody)
	accessToken, authError := ap.SignInWithEmailPassword(loginRequestBody.Email, loginRequestBody.Password)

	if authError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": authError.Error(), "msg": "Login failed - Invalid credentials"})
		return
	}

	user, getUserInfoError := ap.GetUserInfo(accessToken)

	if getUserInfoError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": getUserInfoError.Error(), "msg": "Failed to get user information"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": user.Email, "token": user.AccessToken, "interactionToken": user.InteractionToken})
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
