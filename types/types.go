package types

type LoginRequestBody struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type ApplicationEnv struct {
	SupabaseUrl    string
	SupabaseApiKey string
}
