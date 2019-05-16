package api

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"gopkg.in/resty.v1"
)

func allDocs(url string, data interface{}) error {

	if err := isPointer(data); err != nil {
		return err
	}

	resp, err := r().Get(replaceSpaces(url))

	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return writeRequestError(resp)
	}

	str := string(resp.Body())
	str = str[strings.Index(str, "[") : strings.LastIndex(str, "]")+1]

	return json.Unmarshal([]byte(str), &data)
}

func rowCount(url string) (int, error) {

	type Response struct {
		Rows []struct {
			Key   interface{} `json:"key"`
			Value int         `json:"value"`
		} `json:"rows"`
	}

	resp, err := rr(Response{}).Get(replaceSpaces(url))
	if err != nil {
		return -1, err
	}

	if !resp.IsSuccess() {
		return -1, writeRequestError(resp)
	}

	return (*resp.Result().(*Response)).Rows[0].Value, nil
}

func insert(url string, data interface{}) error {
	resp, err := r().SetBody(data).Post(replaceSpaces(url))
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return writeRequestError(resp)
	}

	return nil
}

func delete(url string) error {
	resp, err := r().Delete(url)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return writeRequestError(resp)
	}

	return nil
}

// rr request with specific response type
func rr(data interface{}) *resty.Request {
	return r().SetResult(data)
}

// encapsulation for resty.R()
func r() *resty.Request {
	return resty.R()
}

// checks if given Type is a Pointer
func isPointer(t interface{}) error {
	if val := reflect.ValueOf(t); val.Kind() == reflect.Ptr {
		return nil
	}
	return fmt.Errorf("Fehler: Der angegebene Typ ist kein Pointer")
}

func docByID(id interface{}, url string, data interface{}) error {
	if err := isPointer(data); err != nil {
		return err
	}

	resp, err := rr(data).Get(replaceSpaces(url))
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		if resp.StatusCode() == 404 {
			return NotFoundError{Msg: "Error: Doc not found"}
		}
		return writeRequestError(resp)
	}

	return nil
}

func replaceSpaces(str string) string {
	return strings.ReplaceAll(str, " ", "%20")
}
