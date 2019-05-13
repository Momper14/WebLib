package users

import (
	"github.com/Momper14/weblib/api"
	"github.com/Momper14/weblib/client"
)

// Users database Users
type Users struct {
	db    api.DB
	views struct {
	}
}

// User struct of a User
type User struct {
	ID       string `json:"_id"`
	Rev      string `json:"_rev"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
	Seit     string `json:"Seit"`
}

// UserByID gibt den User mit der angegebenen ID zurück
func (db Users) UserByID(id string) (User, error) {
	doc := User{}

	if err := db.db.DocByID(id, &doc); err != nil {
		return doc, err
	}

	return doc, nil
}

// AnzahlUsers gibt die Anzahl an User zurück
func (db Users) AnzahlUsers() (int, error) {
	return db.db.RowCount()
}

// New erzeugt einen neuen Users-Handler
func New() Users {
	var db Users

	d := api.New(client.HostURL).DB("users")
	db.db = d

	return db
}
