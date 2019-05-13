package karteikaesten

import (
	"github.com/Momper14/weblib/api"
)

// OeffentlichKastenidKartenid view kastenid-kartenid
type OeffentlichKastenidKartenid struct {
	api.View
}

// OeffentlichKastenidKartenidRow row from view kastenid-kartenid
type OeffentlichKastenidKartenidRow struct {
	ID       string `json:"id"`
	KastenID string `json:"key"`
	KartenID string `json:"value"`
}

// AllDocs returns all Docs from a DB
func (v OeffentlichKastenidKartenid) AllDocs(rows *[]OeffentlichKastenidKartenidRow) error {
	return v.View.AllDocs(rows)
}

// DocsByKey returns all Docs matching the given key
func (v OeffentlichKastenidKartenid) DocsByKey(key string, rows *[]OeffentlichKastenidKartenidRow) error {
	return v.View.DocsByKey(key, rows)
}
