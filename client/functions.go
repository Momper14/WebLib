package client

import "github.com/Momper14/weblib/api"

// GetUUID returns an UUID
func GetUUID() (string, error) {
	return api.New(HostURL).GetUUID()
}
