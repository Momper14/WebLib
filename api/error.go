package api

import (
	"fmt"
	"time"

	"gopkg.in/resty.v1"
)

// RequestError error of a Request
type RequestError struct {
	Code       int
	Status     string
	Time       time.Duration
	ReceivedAt time.Time
	URL        string
	Body       []byte
}

func (e RequestError) Error() string {
	return fmt.Sprintf("Zugfriffsfehler\n\tURL: %s\n\tStatus: %s\n\tBody: %s", e.URL, e.Status, e.Body)
}

func writeRequestError(resp *resty.Response) RequestError {
	return RequestError{
		Code:       resp.StatusCode(),
		Status:     resp.Status(),
		Time:       resp.Time(),
		ReceivedAt: resp.ReceivedAt(),
		URL:        resp.Request.URL,
		Body:       resp.Body(),
	}
}

// NotFoundError somethin is not found
type NotFoundError struct {
	Msg string
}

func (e NotFoundError) Error() string {
	return e.Msg
}
