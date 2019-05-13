package lernen

import (
	"github.com/Momper14/weblib/api"
)

// FachNachKarte view fach-nach-karte
type FachNachKarte struct {
	api.View
}

// FachNachKarteRow row from view fach-nach-karte
type FachNachKarteRow struct {
	ID   string   `json:"id"`
	Key  []string `json:"key"`
	Fach int      `json:"value"`
}

// AllDocs returns all Docs
func (v FachNachKarte) AllDocs(rows *[]FachNachKarteRow) error {
	return v.View.AllDocs(rows)
}

// DocsByKey returns all Docs matching the given key
func (v FachNachKarte) DocsByKey(key string, rows *[]FachNachKarteRow) error {
	return v.View.DocsByKey(key, rows)
}
