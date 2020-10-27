package helper

import "errors"

var (
	ErrRowsAffected = errors.New("could not get affected rows")
	ErrInsertionFailed = errors.New("insertion failed")
	ErrDataInvalid = errors.New("the data entered is not valid")
)
