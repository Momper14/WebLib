package client

import (
	"fmt"
	"log"
	"os"
)

// HostURL Host URL
var HostURL string

func init() {
	pass := os.Getenv("WEBPASS")
	if pass == "" {
		log.Fatal("variable WEBPASS nicht gesetzt!. Mit 'export WEBPASS=pass' setzen.")
	}
	HostURL = fmt.Sprintf("http://admin:%s@localhost:5984", pass)
}
