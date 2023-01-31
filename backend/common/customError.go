package common

// apiMessages contains messages returned by the api

// Error messages
type CustomError struct {
	Message string
}

func (err CustomError) Error() string {
	return err.Message
}
