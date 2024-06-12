package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorverasm/my-community/types"
)

func NewLoginHandler(ap types.AuthProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		handleLogin(c, ap)
	}
}

func NewSignUpHandler(ap types.AuthProvider, cp types.CommunicationProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		handleSignUp(c, ap, cp)
	}
}

func handleLogin(c *gin.Context, ap types.AuthProvider) {
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

	c.JSON(http.StatusOK, gin.H{"email": user.Email, "token": user.AccessToken, "communicationToken": user.CommunicationToken})
}

func handleSignUp(c *gin.Context, ap types.AuthProvider, cp types.CommunicationProvider) {
	var signUpRequestBody types.SignUpRequestBody
	c.BindJSON(&signUpRequestBody)
	communicationToken, createCommunicationTokenError := cp.GetUserCommunicationToken(signUpRequestBody.Email)

	if createCommunicationTokenError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": createCommunicationTokenError.Error(), "msg": "Failed to create token for user"})
	}

	unverifiedUser, signUpError := ap.SignUp(signUpRequestBody.Email, signUpRequestBody.Password, communicationToken)

	if signUpError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": signUpError.Error(), "msg": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": unverifiedUser.Email, "communicationToken": unverifiedUser.CommunicationToken, "msg": "Email confirmation was sent to your email address"})
}
