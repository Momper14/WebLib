package api

import (
	"fmt"
)

const (
	// AllDocs path to all_docs view
	AllDocs = "_design/util/_view/all_docs"
)

// DB DB Handler
type DB struct {
	client Client
	name   string
}

// AllDocs returns all Docs from DB
func (db DB) AllDocs(data interface{}) error {
	url := fmt.Sprintf("%s/%s?reduce=false", db.url(), AllDocs)
	return allDocs(url, data)
}

// RowCount returns number of Rows in DB
func (db DB) RowCount() (int, error) {
	url := fmt.Sprintf("%s/%s", db.url(), AllDocs)
	return rowCount(url)
}

// url returns the URL to the DB
func (db DB) url() string {
	return (fmt.Sprintf("%s/%s", db.client.hostURL, db.name))
}

// DocByID gets a Document by ID
func (db DB) DocByID(id interface{}, data interface{}) error {
	url := fmt.Sprintf("%s/%v", db.url(), id)
	return docByID(id, url, data)
}

// InsertDoc insert the Data into the DB
func (db DB) InsertDoc(data interface{}) error {
	return insert(db.url(), data)
}

// DeleteDoc deletes the Doc with the given id
func (db DB) DeleteDoc(id interface{}) error {
	row, err := db.getRow(id)
	if err != nil {
		return err
	}

	return delete(fmt.Sprintf("%s/%v?rev=%v", db.url(), id, row.Value))
}

// Exists checks if a Doc with the given id exists
func (db DB) Exists(id interface{}) (bool, error) {
	_, err := db.getRow(id)
	if err != nil {
		if _, ok := err.(NotFoundError); ok {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (db DB) getRow(id interface{}) (RowView, error) {
	var data []RowView
	if val, ok := id.(string); ok {
		if val[0] != '[' {
			id = fmt.Sprintf("\"%s\"", val)
		}
	}

	if err := allDocs(fmt.Sprintf("%s/%s?reduce=false&key=%v", db.url(), AllDocs, id), &data); err != nil {
		return RowView{}, err
	}

	if len(data) == 0 {
		return RowView{}, NotFoundError{Msg: "no doc found"}
	}

	return data[0], nil
}

// View returns a View handler
func (db DB) View(design, name string) View {
	return View{
		db:     db,
		design: design,
		name:   name,
	}
}
