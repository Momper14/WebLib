package client

type (
	// NotFoundError someting requested was not found
	NotFoundError struct {
		Msg string
	}

	// ForbiddenError user was not allowed to serve it
	ForbiddenError struct {
		Msg string
	}
)

func (e NotFoundError) Error() string {
	return e.Msg
}

func (e ForbiddenError) Error() string {
	return e.Msg
}
