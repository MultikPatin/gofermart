package schemas

//easyjson -all internal/schemas/auth.go

type AuthCredentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
