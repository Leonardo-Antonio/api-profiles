package user

type iuser interface {
	GetAll() (users []Model, err error)
	SingUp(data Model) error
	SingIn(identification Model) (data Model, err error)
	Delete(identification Model) error
	Update(data Model) error
}
