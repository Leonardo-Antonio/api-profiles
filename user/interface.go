package user

type iuser interface {
	GetAll() (users []model, err error)
	SingUp(data model) error
	SingIn(identification model) (data model, err error)
	Delete(identification model) error
	Update(data model) error
}
