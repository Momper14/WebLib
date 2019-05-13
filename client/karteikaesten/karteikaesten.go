package karteikaesten

import (
	"github.com/Momper14/weblib/api"
	"github.com/Momper14/weblib/client"
)

// Karteikaesten database Karteikaesten
type Karteikaesten struct {
	db    api.DB
	views struct {
		OeffentlichKastenidKartenindex OeffentlichKastenidKartenindex
		NachAutor                      NachAutor
		OeffentlichNachKategorie       OeffentlichNachKategorie
	}
}

// Row row from Karteikaesten
type Row struct {
	ID       string `json:"id"`
	KastenID string `json:"key"`
	Rev      string `json:"value"`
}

// AnzahlOeffentlicherKaesten gibt die Anzahl öffentlicher Karteikästen zurück
func (db Karteikaesten) AnzahlOeffentlicherKaesten() (int, error) {
	return db.views.OeffentlichNachKategorie.RowCount()
}

// AnzahlKaestenUser gibt die Anzahl von Karteikästen zurück,
// welcher der angegebene User erstellt hat
func (db Karteikaesten) AnzahlKaestenUser(id string) (int, error) {
	return db.views.NachAutor.RowCountByKey(id)
}

// AnzahlOeffentlicherKarten gibt die Anzahl öffentlicher Karteikarten zurück
func (db Karteikaesten) AnzahlOeffentlicherKarten() (int, error) {
	return db.views.OeffentlichKastenidKartenindex.RowCount()
}

// OeffentlicheKaestenByKategorie Gibt einen Array mit allen öffentlichen Karteikästen
// der angegebenen Kategorie zurück
func (db Karteikaesten) OeffentlicheKaestenByKategorie(kategorie string) ([]Karteikasten, error) {
	var kaesten []Karteikasten
	rows := []OeffentlichNachKategorieRow{}

	if err := db.views.OeffentlichNachKategorie.DocsByKey(kategorie, &rows); err != nil {
		return kaesten, err
	}

	for _, row := range rows {
		kasten, err := db.KastenByID(row.ID)
		if err != nil {
			return kaesten, err
		}
		kaesten = append(kaesten, kasten)
	}

	return kaesten, nil
}

// KastenByID Gibt den Karteikasten der angegebenen ID zurück
func (db Karteikaesten) KastenByID(id string) (Karteikasten, error) {
	doc := Karteikasten{}

	if err := db.db.DocByID(id, &doc); err != nil {
		return doc, err
	}

	return doc, nil
}

// New erzeugt einen neuen Karteikasten-Handler
func New() Karteikaesten {
	var db Karteikaesten

	d := api.New(client.HostURL).DB("karteikaesten")
	db.db = d

	db.views.OeffentlichKastenidKartenindex = OeffentlichKastenidKartenindex{
		View: d.View("karten", "oeffentlich-kastenid-kartenindex"),
	}

	db.views.NachAutor = NachAutor{
		View: d.View("kasten", "nach-autor"),
	}

	db.views.OeffentlichNachKategorie = OeffentlichNachKategorie{
		View: d.View("kasten", "oeffentlich-nach-kategorie")}

	return db
}
