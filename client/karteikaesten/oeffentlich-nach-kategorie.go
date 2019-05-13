package karteikaesten

import (
	"github.com/Momper14/weblib/api"
)

// OeffentlichNachKategorie view oeffentlich-nach-kategorie
type OeffentlichNachKategorie struct {
	api.View
}

// OeffentlichNachKategorieRow row from view oeffentlich-nach-kategorie
type OeffentlichNachKategorieRow struct {
	ID       string `json:"id"`
	Autor    string `json:"key"`
	KastenID string `json:"value"`
}

// AllDocs returns all Docs
func (v OeffentlichNachKategorie) AllDocs(rows *[]OeffentlichNachKategorieRow) error {
	return v.View.AllDocs(rows)
}

// DocsByKey returns all Docs matching the given key
func (v OeffentlichNachKategorie) DocsByKey(key string, rows *[]OeffentlichNachKategorieRow) error {
	return v.View.DocsByKey(key, rows)
}
