package lernen

import (
	"github.com/Momper14/weblib/api"
)

// KastenNachID view kasten/nach-id
type KastenNachID struct {
	api.View
}

// KastenNachIDRow row from view kasten/nach-id
type KastenNachIDRow struct {
	ID       string `json:"id"`
	KastenID string `json:"key"`
	Rev      string `json:"value"`
}

// AllDocs returns all Docs
func (v KastenNachID) AllDocs(rows *[]KastenNachIDRow) error {
	return v.View.AllDocs(rows)
}

// DocsByKey returns all Docs matching the given key
func (v KastenNachID) DocsByKey(key string, rows *[]KastenNachIDRow) error {
	return v.View.DocsByKey(key, rows)
}
