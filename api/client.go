package api

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
