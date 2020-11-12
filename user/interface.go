package user

type iuser interface {
	GetAll() (users []model, err error)
	SignUp(data model) error
	SignIn(identification model) (data model, err error)
	Delete(identification model) error
	Update(data model) error
}
