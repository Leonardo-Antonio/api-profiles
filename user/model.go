package user

type model struct {
	ID       int    `json:"id" xml:"id"`
	Name     string `json:"name" xml:"name"`
	LastName string `json:"last_name" xml:"last_name"`
	Phone    string `json:"phone" xml:"phone"`
	Email    string `json:"email" xml:"email"`
	Password string `json:"password" xml:"password"`
	Profile  string `json:"profile" xml:"profile"`
}
