package common

// customError contains CustomError struct definition and it's methods

// CustomError allows to use errors with given message. Implements error interface
type CustomError struct {
	Message string
}

func (err CustomError) Error() string {
	return err.Message
}
