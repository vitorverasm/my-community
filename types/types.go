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
	Code  string `json:"code"`
	Email string `json:"email"`
}

type ApplicationEnv struct {
	SupabaseUrl    string
	SupabaseApiKey string
}
