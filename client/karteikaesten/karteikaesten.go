package karteikaesten

import (
	"fmt"

	"github.com/Momper14/weblib/api"
	"github.com/Momper14/weblib/client"
	"github.com/Momper14/weblib/client/lernen"
)

// Karteikaesten database Karteikaesten
type Karteikaesten struct {
	db    api.DB
	views struct {
		OeffentlichKastenidKartenindex OeffentlichKastenidKartenindex
		NachAutor                      NachAutor
		OeffentlichNachKategorie       OeffentlichNachKategorie
		KartenNachAutor                KartenNachAutor
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

// AnzahlKartenUser gibt die Anzahl von Karteikarten zurück,
// welcher der angegebene User erstellt hat
func (db Karteikaesten) AnzahlKartenUser(id string) (int, error) {
	return db.views.KartenNachAutor.RowCountByKey(id)
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

// KastenAnlegen fügt den angegebenen Karteikasten in die Datenbank
func (db Karteikaesten) KastenAnlegen(kasten Karteikasten) error {
	if err := db.db.InsertDoc(kasten); err != nil {
		return err
	}

	lerne := lernen.Lerne{
		User:   kasten.Autor,
		Kasten: kasten.ID,
		Karten: []int{},
	}

	return lernen.New().NeuesLerne(lerne)
}

// KastenBearbeiten aktualisiert den angegebenen Karteikasten in die Datenbank
func (db Karteikaesten) KastenBearbeiten(kasten Karteikasten) error {
	err := db.db.UpdateDoc(kasten.ID, kasten)

	if _, ok := err.(api.NotFoundError); ok {
		return client.NotFoundError{Msg: fmt.Sprintf("Fehler: Kasten %s nicht gefunden", kasten.ID)}
	}

	return err
}

// KastenLoeschen löscht den angegebenen Kasten und alle Lernzustände
func (db Karteikaesten) KastenLoeschen(id string) error {
	// Kasten existiert?
	if _, err := db.KastenByID(id); err != nil {
		return err
	}

	if err := lernen.New().LoeschenAllerLerneZuKasten(id); err != nil {
		return err
	}

	return db.db.DeleteDoc(id)
}

// KastenByID Gibt den Karteikasten der angegebenen ID zurück
func (db Karteikaesten) KastenByID(id string) (Karteikasten, error) {
	doc := Karteikasten{}

	if err := db.db.DocByID(id, &doc); err != nil {
		if _, ok := err.(api.NotFoundError); ok {
			return doc, client.NotFoundError{Msg: fmt.Sprintf("Fehler: Kasten %s nicht gefunden", id)}
		}
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
		View: d.View("kasten", "oeffentlich-nach-kategorie"),
	}

	db.views.KartenNachAutor = KartenNachAutor{
		View: d.View("karten", "nach-autor"),
	}

	return db
}
