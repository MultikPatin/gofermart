package dtos

type AuthCredentials struct {
	Login    string
	Password string
}

type User struct {
	ID       int64
	Login    string
	Password string
}
