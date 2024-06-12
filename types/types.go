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

type ApplicationEnv struct {
	Environment     string
	SupabaseUrl     string
	SupabaseApiKey  string
	StreamApiKey    string
	StreamApiSecret string
}
