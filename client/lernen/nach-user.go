package lernen

import (
	"github.com/Momper14/weblib/api"
)

// NachUser view kastenid-kartenid
type NachUser struct {
	api.View
}

// NachUserRow row from view kastenid-kartenid
type NachUserRow struct {
	ID   string `json:"id"`
	User string `json:"key"`
	Rev  string `json:"value"`
}

// AllDocs returns all Docs from a DB
func (v NachUser) AllDocs(rows *[]NachUserRow) error {
	return v.View.AllDocs(rows)
}

// DocsByKey returns all Docs matching the given key
func (v NachUser) DocsByKey(key string, rows *[]NachUserRow) error {
	return v.View.DocsByKey(key, rows)
}
