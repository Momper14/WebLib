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

// View returns a View handler
func (db DB) View(design, name string) View {
	return View{
		db:     db,
		design: design,
		name:   name,
	}
}
