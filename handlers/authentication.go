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

	c.JSON(http.StatusOK, gin.H{"email": user.Email, "token": user.AccessToken, "interactionToken": user.InteractionToken})
}

func handleSignUp(c *gin.Context, ap types.AuthProvider, cp types.CommunicationProvider) {
	var signUpRequestBody types.SignUpRequestBody
	c.BindJSON(&signUpRequestBody)
	interactionToken, createInteractionTokenError := cp.GetUserInteractionToken(signUpRequestBody.Email)

	if createInteractionTokenError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": createInteractionTokenError.Error(), "msg": "Failed to create token for user"})
	}

	unverifiedUser, signUpError := ap.SignUp(signUpRequestBody.Email, signUpRequestBody.Password, interactionToken)

	if signUpError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": signUpError.Error(), "msg": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": unverifiedUser.Email, "interactionToken": unverifiedUser.InteractionToken, "msg": "Email confirmation was sent to your email address"})
}
