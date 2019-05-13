package karteikaesten

import (
	"github.com/Momper14/weblib/api"
)

// OeffentlichKastenidKartenindex view oeffentlich-Kastenid-kartenindex
type OeffentlichKastenidKartenindex struct {
	api.View
}

// OeffentlichKastenidKartenindexRow row from view oeffentlich-Kastenid-kartenindex
type OeffentlichKastenidKartenindexRow struct {
	ID          string `json:"index"`
	KastenID    string `json:"key"`
	KartenIndex string `json:"value"`
}

// AllDocs returns all Docs
func (v OeffentlichKastenidKartenindex) AllDocs(rows *[]OeffentlichKastenidKartenindexRow) error {
	return v.View.AllDocs(rows)
}

// DocsByKey returns all Docs matching the given key
func (v OeffentlichKastenidKartenindex) DocsByKey(key string, rows *[]OeffentlichKastenidKartenindexRow) error {
	return v.View.DocsByKey(key, rows)
}
