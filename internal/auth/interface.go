package auth

type AuthRepository interface {
	GetLoginMethod() string
}

type AuthService interface {
	GetLoginMethod() string
}
