package users

import (
	"github.com/Momper14/weblib/api"
)

// NachEMail view nach-user
type NachEMail struct {
	api.View
}

// NachEMailRow row from nach-user
type NachEMailRow struct {
	ID    string `json:"id"`
	EMail string `json:"key"`
	Val   string `json:"value"`
}

// AllDocs returns all Docs
func (v NachEMail) AllDocs(rows *[]NachEMailRow) error {
	return v.View.AllDocs(rows)
}

// DocsByKey returns all Docs matching the given key
func (v NachEMail) DocsByKey(key string, rows *[]NachEMailRow) error {
	return v.View.DocsByKey(key, rows)
}
