package api

import (
	"fmt"
)

// View DB Handler
type View struct {
	db     DB
	design string
	name   string
}

// RowView struct for a Row of a View
type RowView struct {
	ID    Value `json:"id"`
	Key   Value `json:"key"`
	Value Value `json:"value"`
}

// AllDocs returns all Docs from a DB
func (v View) AllDocs(data Value) error {
	return allDocs(fmt.Sprintf("%s?reduce=false", v.url()), data)
}

// DocsByKey returns all Docs matching the given key
func (v View) DocsByKey(key Value, data Value) error {
	if val, ok := key.(string); ok {
		if val[0] != '[' {
			key = fmt.Sprintf("\"%s\"", val)
		}
	}
	return allDocs(fmt.Sprintf("%s?reduce=false&key=%v", v.url(), key), data)
}

// RowCount returns count of Rows
func (v View) RowCount() (int, error) {
	return rowCount(v.url())
}

// RowCountByKey returns count of Rows
func (v View) RowCountByKey(key Value) (int, error) {
	return rowCount(fmt.Sprintf("%s?key=\"%s\"", v.url(), key))
}

func (v View) url() string {
	return (fmt.Sprintf("%s/_design/%s/_view/%s", v.db.url(), v.design, v.name))
}
