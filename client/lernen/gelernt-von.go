package lernen

import (
	"github.com/Momper14/weblib/api"
)

// GelerntVon view kastenid-kartenid
type GelerntVon struct {
	api.View
}

// GelerntVonRow row from view kastenid-kartenid
type GelerntVonRow struct {
	ID  string   `json:"id"`
	Key []string `json:"key"`
	Rev string   `json:"value"`
}

// AllDocs returns all Docs from a DB
func (v GelerntVon) AllDocs(rows *[]GelerntVonRow) error {
	return v.View.AllDocs(rows)
}

// DocsByKey returns all Docs matching the given key
func (v GelerntVon) DocsByKey(key string, rows *[]GelerntVonRow) error {
	return v.View.DocsByKey(key, rows)
}
