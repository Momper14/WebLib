package kategorien

import (
	"github.com/Momper14/weblib/api"
	"github.com/Momper14/weblib/client"
)

// Kategorien database Kategorien
type Kategorien struct {
	db api.DB
	//views struct {}
}

// Kategorie struct of a Karteikasten
type Kategorie struct {
	ID             string   `json:"_id"`
	Rev            string   `json:"_rev"`
	Unterkategorie []string `json:"Unterkategorie"`
}

// AlleKategorien gibt einen Array mit allen Kategorien zurück
func (db Kategorien) AlleKategorien() ([]Kategorie, error) {
	var kategorien []Kategorie
	rows := []api.RowView{}

	if err := db.db.AllDocs(&rows); err != nil {
		return nil, err
	}

	for _, row := range rows {
		kategorie, err := db.kategorieByID(row.ID)
		if err != nil {
			return nil, err
		}
		kategorien = append(kategorien, kategorie)
	}

	return kategorien, nil
}

// kategorieByID gibt die Kategorie mit der angegebenen ID zurück
func (db Kategorien) kategorieByID(id string) (Kategorie, error) {
	doc := Kategorie{}

	if err := db.db.DocByID(id, &doc); err != nil {
		return doc, err
	}

	return doc, nil
}

// New erzeugt einen neuen Kategorien-Handler
func New() Kategorien {
	var db Kategorien

	d := api.New(client.HostURL).DB("kategorien")
	db.db = d

	return db
}
