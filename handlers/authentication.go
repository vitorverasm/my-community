package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gotrueTypes "github.com/supabase-community/gotrue-go/types"
	supabaseSdk "github.com/supabase-community/supabase-go"
	"github.com/vitorverasm/my-community/types"
)

func HandleLogin(c *gin.Context, sp *supabaseSdk.Client) {
	var loginRequestBody types.LoginRequestBody
	c.BindJSON(&loginRequestBody)
	token, error := sp.Auth.SignInWithEmailPassword(loginRequestBody.Email, loginRequestBody.Password)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": loginRequestBody.Email, "token": token.AccessToken})
}

func HandleSignUp(c *gin.Context, sp *supabaseSdk.Client) {
	var signUpRequestBody types.SignUpRequestBody
	c.BindJSON(&signUpRequestBody)

	res, error := sp.Auth.Signup(gotrueTypes.SignupRequest{
		Email:    signUpRequestBody.Email,
		Password: signUpRequestBody.Password,
	})
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": res.User.Email, "token": res.AccessToken, "email_confirmed_at": res.User.EmailConfirmedAt})
}

func HandleMagicLink(c *gin.Context, sp *supabaseSdk.Client) {
	var magicLinkRequestBody types.MagicLinkRequestBody
	c.BindJSON(&magicLinkRequestBody)
	error := sp.Auth.OTP(gotrueTypes.OTPRequest{
		Email:      magicLinkRequestBody.Email,
		CreateUser: true,
	})

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": magicLinkRequestBody.Email, "message": "Sign up link sent to E-mail"})
}

func HandleValidateOTP(c *gin.Context, sp *supabaseSdk.Client) {
	var validateOTPRequestBody types.ValidateOTPRequestBody
	c.BindJSON(&validateOTPRequestBody)
	res, error := sp.Auth.VerifyForUser(gotrueTypes.VerifyForUserRequest{
		Type:       "signup",
		Token:      validateOTPRequestBody.Code,
		Email:      validateOTPRequestBody.Email,
		RedirectTo: "http://localhost:3000",
	})

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": res.User.Email, "token": res.AccessToken})
}
