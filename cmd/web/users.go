package main

import (
	"net/http"
)

func (app *Application) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "home.gohtml", data)
}

func (app *Application) userListHandler(w http.ResponseWriter, r *http.Request) {
	users, err := app.Models.UserModel.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Users = users

	app.render(w, http.StatusOK, "user_list.gohtml", data)
}

func (app *Application) userCreateHandler(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "user_create.gohtml", data)
}

func (app *Application) userDetailHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		app.notFound(w)
		return
	}

	user, err := app.Models.UserModel.GetById(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.User = user

	app.render(w, http.StatusOK, "user_detail.gohtml", data)
}

func (app *Application) userUpdateHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		app.notFound(w)
		return
	}

	user, err := app.Models.UserModel.GetById(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.User = user

	app.redirect(w, r, "/users")
}

func (app *Application) userDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		app.notFound(w)
		return
	}

	user, err := app.Models.UserModel.GetById(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = app.Models.UserModel.Delete(user.ID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.redirect(w, r, "/users")
}
