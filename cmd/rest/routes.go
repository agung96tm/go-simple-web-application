package main

import "net/http"

func (app *Application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", app.userListHandler)
	mux.HandleFunc("/user-detail", app.userDetailHandler)
	mux.HandleFunc("/user-create", app.userCreateHandler)
	mux.HandleFunc("/user-update", app.userUpdateHandler)
	mux.HandleFunc("/user-delete", app.userDeleteHandler)

	return mux
}
