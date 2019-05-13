package karteikaesten

import (
	"github.com/Momper14/weblib/api"
	"github.com/Momper14/weblib/client"
)

// Karteikaesten database Karteikaesten
type Karteikaesten struct {
	db    api.DB
	views struct {
		OeffentlichKastenindexKartenindex OeffentlichKastenindexKartenindex
		NachAutor                         NachAutor
		OeffentlichNachKategorie          OeffentlichNachKategorie
	}
}

// Row row from view kastenid-kartenid
type Row struct {
	ID       string `json:"id"`
	KastenID string `json:"key"`
	Rev      string `json:"value"`
}

// AnzahlOeffentlicherKaesten returns count of Rows
func (db Karteikaesten) AnzahlOeffentlicherKaesten() (int, error) {
	return db.views.OeffentlichNachKategorie.RowCount()
}

// AnzahlKaestenUser returns count of Rows
func (db Karteikaesten) AnzahlKaestenUser(id string) (int, error) {
	return db.views.NachAutor.RowCountByKey(id)
}

// AnzahlOeffentlicherKarten returns count of Rows
func (db Karteikaesten) AnzahlOeffentlicherKarten() (int, error) {
	return db.views.OeffentlichKastenindexKartenindex.RowCount()
}

// OeffentlicheKaestenByKategorie returns a list with all Karteikasten of given Kategorie
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

// KastenByID returns all Docs matching the given key
func (db Karteikaesten) KastenByID(id string) (Karteikasten, error) {
	doc := Karteikasten{}

	if err := db.db.DocByID(id, &doc); err != nil {
		return doc, err
	}

	return doc, nil
}

// New creates a Karteikaesten
func New() Karteikaesten {
	var db Karteikaesten

	d := api.New(client.HostURL).DB("karteikaesten")
	db.db = d

	db.views.OeffentlichKastenindexKartenindex = OeffentlichKastenindexKartenindex{
		View: d.View("karten", "oeffentlich-kastenid-kartenindex"),
	}

	db.views.NachAutor = NachAutor{
		View: d.View("kasten", "nach-autor"),
	}

	db.views.OeffentlichNachKategorie = OeffentlichNachKategorie{
		View: d.View("kasten", "oeffentlich-nach-kategorie")}

	return db
}
