package helper

import "errors"

var (
	ErrRowsAffected = errors.New("could not get affected rows")
	ErrInsertionFailed = errors.New("insertion failed")
	ErrDataInvalid = errors.New("the data entered is not valid")
	ErrEmailInvalid = errors.New("the email entered is invalid")
	ErrPasswordNotSecure = errors.New("the password entered is not secure")
	ErrNickNameIvalid = errors.New("the nick name is invalid")
	ErrRowNotAffected = errors.New("rows were not affected")
)

var (
	ERROR = "error"
	MESSAGE = "message"
)