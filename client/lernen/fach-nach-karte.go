package lernen

import (
	"github.com/Momper14/weblib/api"
)

// FachNachKarte view kastenid-kartenid
type FachNachKarte struct {
	api.View
}

// FachNachKarteRow row from view kastenid-kartenid
type FachNachKarteRow struct {
	ID       string   `json:"id"`
	KartenID []string `json:"key"`
	Fach     int      `json:"value"`
}

// AllDocs returns all Docs from a DB
func (v FachNachKarte) AllDocs(rows *[]FachNachKarteRow) error {
	return v.View.AllDocs(rows)
}

// DocsByKey returns all Docs matching the given key
func (v FachNachKarte) DocsByKey(key string, rows *[]FachNachKarteRow) error {
	return v.View.DocsByKey(key, rows)
}
