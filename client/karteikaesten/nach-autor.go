package karteikaesten

import (
	"github.com/Momper14/weblib/api"
)

// NachAutor view nach-autor
type NachAutor struct {
	api.View
}

// NachAutorRow row from view nach-autor
type NachAutorRow struct {
	ID       string `json:"id"`
	Autor    string `json:"key"`
	KastenID string `json:"value"`
}

// AllDocs returns all Docs
func (v NachAutor) AllDocs(rows *[]NachAutorRow) error {
	return v.View.AllDocs(rows)
}

// DocsByKey returns all Docs matching the given key
func (v NachAutor) DocsByKey(key string, rows *[]NachAutorRow) error {
	return v.View.DocsByKey(key, rows)
}
