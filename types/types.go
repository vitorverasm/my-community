package types

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MagicLinkRequestBody struct {
	Email string `json:"email"`
}

type ValidateOTPRequestBody struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type User struct {
	Email            string `json:"email"`
	InteractionToken string `json:"interactionToken"`
	AccessToken      string `json:"token"`
}

type UnverifiedUser struct {
	Email            string
	InteractionToken string
}

type AuthProvider interface {
	SignInWithEmailPassword(email string, password string) (accessToken string, err error)
	GetUserInfo(accessToken string) (User, error)
	SignUp(email string, password string, interactionToken string) (UnverifiedUser, error)
}

type ApplicationEnv struct {
	Environment     string
	SupabaseUrl     string
	SupabaseApiKey  string
	StreamApiKey    string
	StreamApiSecret string
}
