package lernen

import (
	"fmt"

	"github.com/Momper14/weblib/api"
	"github.com/Momper14/weblib/client"
)

// Lernen database Lernen
type Lernen struct {
	db    api.DB
	views struct {
		GelerntVon    GelerntVon
		NachUser      NachUser
		FachNachKarte FachNachKarte
	}
}

// Lerne struct of a Karteikasten
type Lerne struct {
	ID     string `json:"_id"`
	Rev    string `json:"_rev"`
	User   string `json:"User"`
	Kasten string `json:"Kasten"`
	Karten []int  `json:"Karten"`
}

// LerneByID returns all Docs matching the given key
func (db Lernen) LerneByID(id string) (Lerne, error) {
	doc := Lerne{}

	if err := db.db.DocByID(id, &doc); err != nil {
		return doc, err
	}

	return doc, nil
}

// LerneByUserAndKasten returns all Docs matching the given key
func (db Lernen) LerneByUserAndKasten(userid, kastenid string) (Lerne, error) {
	rows := []GelerntVonRow{}
	lerne := Lerne{}
	key := fmt.Sprintf("[\"%s\", \"%s\"]", userid, kastenid)

	if err := db.views.GelerntVon.DocsByKey(key, &rows); err != nil {
		return lerne, err
	}

	if len(rows) == 0 {
		return lerne, client.NotFoundError{Msg: fmt.Sprintf("Error: User %s hat Kasten %s nicht gelernt", userid, kastenid)}
	}

	lerne, err := db.LerneByID(rows[0].ID)

	return lerne, err
}

// GelerntVonUser returns all Docs matching the given key
func (db Lernen) GelerntVonUser(userid string) ([]Lerne, error) {
	rows := []NachUserRow{}
	var gelerntVon []Lerne

	if err := db.views.NachUser.DocsByKey(userid, &rows); err != nil {
		return gelerntVon, err
	}

	for _, row := range rows {
		lerne, err := db.LerneByID(row.ID)
		if err != nil {
			return gelerntVon, err
		}
		gelerntVon = append(gelerntVon, lerne)
	}
	return gelerntVon, nil
}

// FachVonKarte returns all Docs matching the given key
func (db Lernen) FachVonKarte(userid, kastenid, kartenindex string) (int, error) {
	rows := []FachNachKarteRow{}
	key := fmt.Sprintf("[\"%s\", \"%s\", \"%s\"]", userid, kastenid, kartenindex)

	if err := db.views.FachNachKarte.DocsByKey(key, &rows); err != nil {
		return -1, err
	}
	if len(rows) == 0 {
		return -1, client.NotFoundError{Msg: fmt.Sprintf("Error: Keine karte grfunden")}
	}
	return rows[0].Fach, nil
}

// New creates a Lernen
func New() Lernen {
	var db Lernen

	d := api.New(client.HostURL).DB("lernen")
	db.db = d

	db.views.GelerntVon = GelerntVon{
		View: d.View("kasten", "gelernt-von")}

	db.views.NachUser = NachUser{
		View: d.View("user", "nach-user")}

	db.views.FachNachKarte = FachNachKarte{
		View: d.View("karten", "fach-nach-karte")}

	return db
}
