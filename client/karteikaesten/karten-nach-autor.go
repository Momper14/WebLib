package karteikaesten

import (
	"github.com/Momper14/weblib/api"
)

// KartenNachAutor view nach-autor
type KartenNachAutor struct {
	api.View
}

// KartenNachAutorRow row from view nach-autor
type KartenNachAutorRow struct {
	ID       string `json:"id"`
	Autor    string `json:"key"`
	KartenID string `json:"value"`
}

// AllDocs returns all Docs
func (v KartenNachAutor) AllDocs(rows *[]KartenNachAutorRow) error {
	return v.View.AllDocs(rows)
}

// DocsByKey returns all Docs matching the given key
func (v KartenNachAutor) DocsByKey(key string, rows *[]KartenNachAutorRow) error {
	return v.View.DocsByKey(key, rows)
}
