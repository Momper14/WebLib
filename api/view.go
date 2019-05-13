package api

import (
	"fmt"
)

// View View Handler
type View struct {
	db     DB
	design string
	name   string
}

// RowView struct for a Row of a View
type RowView struct {
	ID    interface{} `json:"id"`
	Key   interface{} `json:"key"`
	Value interface{} `json:"value"`
}

// AllDocs returns all Docs from View
func (v View) AllDocs(data interface{}) error {
	return allDocs(fmt.Sprintf("%s?reduce=false", v.url()), data)
}

// DocsByKey returns all Docs matching the given key
func (v View) DocsByKey(key interface{}, data interface{}) error {
	if val, ok := key.(string); ok {
		if val[0] != '[' {
			key = fmt.Sprintf("\"%s\"", val)
		}
	}
	return allDocs(fmt.Sprintf("%s?reduce=false&key=%v", v.url(), key), data)
}

// RowCount returns number of Rows in View
func (v View) RowCount() (int, error) {
	return rowCount(v.url())
}

// RowCountByKey returns number of Rows in View with given key
func (v View) RowCountByKey(key interface{}) (int, error) {
	return rowCount(fmt.Sprintf("%s?key=\"%s\"", v.url(), key))
}

// url returns the URL to the View
func (v View) url() string {
	return (fmt.Sprintf("%s/_design/%s/_view/%s", v.db.url(), v.design, v.name))
}
