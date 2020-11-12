package helper

import "errors"

var (
	ErrEmailInvalid               = errors.New("the email entered is invalid")
	ErrPersonalInformationInvalid = errors.New("personal information invalid")
	ErrInsecurePassword           = errors.New("the password entered insecure")
	ErrUserInvalid                = errors.New("the user not exist")
	ErrRowNotAffected             = errors.New("the row not affected")
)
