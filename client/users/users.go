package users

import (
	"fmt"

	"github.com/Momper14/weblib/api"
	"github.com/Momper14/weblib/client"
	"github.com/Momper14/weblib/client/karteikaesten"
)

// Users database Users
type Users struct {
	db    api.DB
	views struct {
		NachEMail NachEMail
	}
}

// User struct of a User
type User struct {
	Name     string `json:"_id"`
	Rev      string `json:"_rev,omitempty"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
	Seit     int64  `json:"Seit"`
	Bild     string `json:"Bild"`
}

// UserByID gibt den User mit der angegebenen ID zurück
func (db Users) UserByID(name string) (User, error) {
	doc := User{}

	err := db.db.DocByID(name, &doc)

	if _, ok := err.(api.NotFoundError); ok {
		err = client.NotFoundError{Msg: fmt.Sprintf("Fehler: User %s nicht gefunden", name)}
	}

	return doc, err
}

// AnzahlUsers gibt die Anzahl an User zurück
func (db Users) AnzahlUsers() (int, error) {
	return db.db.RowCount()
}

// UserErstellen fügt den gegebenen User in die Datenbank ein
func (db Users) UserErstellen(user User) error {
	return db.db.InsertDoc(user)
}

// UserAktualisieren aktualisiert den User
func (db Users) UserAktualisieren(user User) error {
	err := db.db.UpdateDoc(user.Name, user)

	if _, ok := err.(api.NotFoundError); ok {
		err = client.NotFoundError{Msg: fmt.Sprintf("Fehler: User %s nicht gefunden", user.Name)}
	}

	return err
}

// UserLoeschen löscht den User mit dem gegebenen Namen
func (db Users) UserLoeschen(name string) error {
	err := db.db.DeleteDoc(name)

	karteikaesten := karteikaesten.New()

	rows, err := karteikaesten.KastenVonUser(name)
	if err != nil {
		return err
	}

	for _, row := range rows {
		if err := karteikaesten.KastenLoeschen(row.KastenID); err != nil {
			return err
		}
	}

	if _, ok := err.(api.NotFoundError); ok {
		err = client.NotFoundError{Msg: fmt.Sprintf("Fehler: User %s nicht gefunden", name)}
	}

	return err
}

// CheckEmail überprüft ob die gegebene E-Mail Addresse vergeben ist
func (db Users) CheckEmail(email string) (bool, error) {
	var data []NachEMailRow

	if err := db.views.NachEMail.DocsByKey(email, &data); err != nil {
		return false, err
	}

	if len(data) > 0 {
		return true, nil
	}

	return false, nil
}

// CheckName überprüft ob der gegebene Name bereits vergeben ist
func (db Users) CheckName(name string) (bool, error) {
	var user User

	if err := db.db.DocByID(name, &user); err != nil {
		if _, ok := err.(api.NotFoundError); ok {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// New erzeugt einen neuen Users-Handler
func New() Users {
	var db Users

	d := api.New(client.HostURL).DB("users")
	db.db = d

	db.views.NachEMail = NachEMail{
		View: d.View("util", "nach-email")}

	return db
}
