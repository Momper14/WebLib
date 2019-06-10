package client

import (
	"log"
	"os"
)

// HostURL Host URL
var HostURL string

func init() {
	HostURL = os.Getenv("COUCHDB_URL")
	if HostURL == "" {
		log.Fatal("variable COUCHDB_URL nicht gesetzt!. Mit 'export COUCHDB_URL=url' setzen.")
	}
}
