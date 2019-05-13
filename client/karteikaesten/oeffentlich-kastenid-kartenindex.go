package karteikaesten

import (
	"github.com/Momper14/weblib/api"
)

// OeffentlichKastenindexKartenindex view kastenindex-kartenindex
type OeffentlichKastenindexKartenindex struct {
	api.View
}

// OeffentlichKastenindexKartenindexRow row from view kastenindex-kartenindex
type OeffentlichKastenindexKartenindexRow struct {
	ID          string `json:"index"`
	KastenID    string `json:"key"`
	KartenIndex string `json:"value"`
}

// AllDocs returns all Docs from a DB
func (v OeffentlichKastenindexKartenindex) AllDocs(rows *[]OeffentlichKastenindexKartenindexRow) error {
	return v.View.AllDocs(rows)
}

// DocsByKey returns all Docs matching the given key
func (v OeffentlichKastenindexKartenindex) DocsByKey(key string, rows *[]OeffentlichKastenindexKartenindexRow) error {
	return v.View.DocsByKey(key, rows)
}
