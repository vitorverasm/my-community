package supabase

import (
	"log"

	gotrueTypes "github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
	"github.com/vitorverasm/my-community/types"
)

func InitializeClient(url string, apiKey string) *supabase.Client {
	client, err := supabase.NewClient(url, apiKey, nil)
	if err != nil {
		log.Println("cannot initialize supabase client", err)
	}

	return client
}

type SupabaseAuthProvider struct {
	Client *supabase.Client
}

func (sup *SupabaseAuthProvider) SignInWithEmailPassword(email string, password string) (accessToken string, err error) {
	token, signInError := sup.Client.Auth.SignInWithEmailPassword(email, password)

	if signInError != nil {
		return "", signInError
	}

	return token.AccessToken, nil
}

func (sup *SupabaseAuthProvider) GetUserInfo(accessToken string) (types.User, error) {
	authorizedClient := sup.Client.Auth.WithToken(
		accessToken,
	)

	user, getUserError := authorizedClient.GetUser()

	if getUserError != nil {
		return types.User{}, getUserError
	}

	interactionToken, ok := user.UserMetadata["interactionToken"].(string)
	if !ok {
		return types.User{
			Email:            user.Email,
			AccessToken:      accessToken,
			InteractionToken: "",
		}, nil
	}

	return types.User{
		Email:            user.Email,
		AccessToken:      accessToken,
		InteractionToken: interactionToken,
	}, nil
}

func (sup *SupabaseAuthProvider) SignUp(email string, password string, interactionToken string) (types.UnverifiedUser, error) {
	res, error := sup.Client.Auth.Signup(gotrueTypes.SignupRequest{
		Email:    email,
		Password: password,
		Data: map[string]interface{}{
			"interactionToken": interactionToken,
		},
	})

	if error != nil {
		return types.UnverifiedUser{}, error
	}

	return types.UnverifiedUser{
		Email:            res.Email,
		InteractionToken: interactionToken,
	}, nil
}
