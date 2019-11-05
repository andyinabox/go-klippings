package types

// User a simple user model that takes only a
// password
type User struct {
	Password string
}

type UserService interface {
	Login() error
	Logout() error
}