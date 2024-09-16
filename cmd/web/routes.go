package main

import "net/http"

func (app *Application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./internal/ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", app.homeHandler)
	mux.HandleFunc("/users", app.userListHandler)
	mux.HandleFunc("/users/create", app.userCreateHandler)
	mux.HandleFunc("/users/view", app.userDetailHandler)
	mux.HandleFunc("/users/update", app.userUpdateHandler)
	mux.HandleFunc("/users/delete", app.userDeleteHandler)

	return mux
}
