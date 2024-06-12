package types

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Email              string `json:"email"`
	CommunicationToken string `json:"communicationToken"`
	AccessToken        string `json:"token"`
}

type UnverifiedUser struct {
	Email              string `json:"email"`
	CommunicationToken string `json:"communicationToken"`
}

type AuthProvider interface {
	SignInWithEmailPassword(email string, password string) (accessToken string, err error)
	GetUserInfo(accessToken string) (User, error)
	SignUp(email string, password string, communicationToken string) (UnverifiedUser, error)
}

type CommunicationProvider interface {
	GetUserCommunicationToken(userEmail string) (communicationToken string, err error)
}

type ApplicationEnv struct {
	Environment     string
	SupabaseUrl     string
	SupabaseApiKey  string
	StreamApiKey    string
	StreamApiSecret string
}
