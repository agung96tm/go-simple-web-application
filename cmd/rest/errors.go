package main

import "net/http"

func (app *Application) errorResponse(w http.ResponseWriter, status int, message interface{}) {
	err := app.writeJSON(w, status, map[string]interface{}{
		"message": message,
	}, nil)

	if err != nil {
		app.Logger.Printf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) notFoundResponse(w http.ResponseWriter) {
	message := "Not found"
	app.errorResponse(w, http.StatusNotFound, message)
}

func (app *Application) badRequestResponse(w http.ResponseWriter, err error) {
	app.errorResponse(w, http.StatusBadRequest, err.Error())
}

func (app *Application) serverErrorResponse(w http.ResponseWriter, err error) {
	app.Logger.Println(err.Error())
	message := "Server error"
	app.errorResponse(w, http.StatusInternalServerError, message)
}
