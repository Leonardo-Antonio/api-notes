package user

type iuser interface {
	CreateUser(user model) error
	Login(data model) (user model, err error)
}