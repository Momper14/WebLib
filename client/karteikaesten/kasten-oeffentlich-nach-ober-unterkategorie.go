package karteikaesten

import "github.com/Momper14/weblib/api"

// OeffentlichNachOberUnterkategorie view oeffentlich-nach-ober-unterkategorie
type OeffentlichNachOberUnterkategorie struct {
	api.View
}

// OeffentlichNachOberUnterkategorieRow row from view oeffentlich-nach-ober-unterkategorie
type OeffentlichNachOberUnterkategorieRow struct {
	ID  string   `json:"id"`
	Key []string `json:"key"`
	Rev string   `json:"value"`
}

// AllDocs returns all Docs
func (v OeffentlichNachOberUnterkategorie) AllDocs(rows *[]OeffentlichNachOberUnterkategorieRow) error {
	return v.View.AllDocs(rows)
}

// DocsByKey returns all Docs matching the given key
func (v OeffentlichNachOberUnterkategorie) DocsByKey(key string, rows *[]OeffentlichNachOberUnterkategorieRow) error {
	return v.View.DocsByKey(key, rows)
}
