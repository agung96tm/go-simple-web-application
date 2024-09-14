package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

func (app *Application) readIDParam(r *http.Request) (string, error) {
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		return "", errors.New("missing required parameter 'id'")
	}
	return idParam, nil
}

func (app *Application) readJSON(r *http.Request, v interface{}) error {
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(v); err != nil {
		return err
	}
	return nil
}

func (app *Application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(js)
	return nil
}
