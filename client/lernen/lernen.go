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

// Lerne struct of a "Lern-state"
type Lerne struct {
	ID     string `json:"_id,omitempty"`
	Rev    string `json:"_rev,omitempty"`
	User   string `json:"User"`
	Kasten string `json:"Kasten"`
	Karten []int  `json:"Karten"`
}

// LerneByID gibt den Lernfortschritt mit der angegebenen ID zurück
func (db Lernen) LerneByID(id string) (Lerne, error) {
	doc := Lerne{}

	if err := db.db.DocByID(id, &doc); err != nil {
		return doc, err
	}

	return doc, nil
}

// LerneByUserAndKasten gibt den Lernfortschritt
// des Users für den Karteikasten zurück
func (db Lernen) LerneByUserAndKasten(userid, kastenid string) (Lerne, error) {
	rows := []GelerntVonRow{}
	lerne := Lerne{}
	key := fmt.Sprintf("[\"%s\", \"%s\"]", userid, kastenid)

	if err := db.views.GelerntVon.DocsByKey(key, &rows); err != nil {
		return lerne, err
	}

	if len(rows) == 0 {
		return lerne, client.NotFoundError{
			Msg: fmt.Sprintf("Error: User %s hat Kasten %s nicht gelernt", userid, kastenid),
		}
	}

	lerne, err := db.LerneByID(rows[0].ID)

	return lerne, err
}

// NeuesLerne trägt einen neuen Lern-Status in die Datenbank
func (db Lernen) NeuesLerne(lerne Lerne) error {
	return db.db.InsertDoc(lerne)
}

// GelerntVonUser gibt alle Lernfortschritte des Users zurück
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

// KarteGelernt karte mit gegebenem index wurde von gegebenem user gelernt
func (db Lernen) KarteGelernt(userid, kastenid string, index int, erfolg bool) error {
	lerne, err := db.LerneByUserAndKasten(userid, kastenid)
	if err != nil {
		return err
	}

	if erfolg {
		if lerne.Karten[index] < 4 {
			lerne.Karten[index]++
		}
	} else {
		lerne.Karten[index] = 0
	}
	return db.AktualisiereLerne(lerne)
}

// AktualisiereLerne speichert änderungen in die Datenbank
func (db Lernen) AktualisiereLerne(lerne Lerne) error {
	return db.db.InsertDoc(lerne)
}

// FachVonKarte gibt das Fach der Karteikarte aus dem Karteikasten für den User zurück
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

// New erzeugt einen neuen Lernen-Handler
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
