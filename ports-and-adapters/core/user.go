package core

type User struct {
	Id string
	Username string
	Email string
}

type UserRepository interface {
	GetUserById(id string) (*User, error)
}