package client

// NotFoundError someting requested was not found
type NotFoundError struct {
	Msg string
}

func (e NotFoundError) Error() string {
	return e.Msg
}
