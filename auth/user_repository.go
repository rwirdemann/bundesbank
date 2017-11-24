package auth

type User struct {
	Username string
	Password string
}

type UserRepository interface {
	FindUser(username string) (User, error)
}
