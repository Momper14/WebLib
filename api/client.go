package api

import (
	"fmt"
)

// Client client for the DB
type Client struct {
	hostURL string
}

// New constructs a new client
func New(hostURL string) Client {
	return Client{
		hostURL: hostURL,
	}
}

// DB creates a BD handler
func (c Client) DB(name string) DB {
	return DB{
		client: c,
		name:   name,
	}
}

// GetUUID returns an UUID
func (c Client) GetUUID() (string, error) {
	type UUIDs struct {
		Uuids []string `json:"uuids"`
	}

	url := fmt.Sprintf("%s/%s", c.hostURL, "/_uuids?count=1")
	var uuids UUIDs

	resp, err := rr(&uuids).Get(url)
	if err != nil {
		return "", err
	}

	if !resp.IsSuccess() {
		return "", writeRequestError(resp)
	}

	return uuids.Uuids[0], nil
}
