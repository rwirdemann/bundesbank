package user

type UserService struct {
	Repository UserRepository
}

func (service UserService) Authenticate(username string, password string) bool {
	if user, err := service.Repository.FindUser(username); err == nil && user.Password == password {
		return true
	}
	return false
}
