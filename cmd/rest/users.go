package main

import (
	"github.com/agung96tm/go-simple-web-application/internal/data"
	"net/http"
)

func (app *Application) userListHandler(w http.ResponseWriter, _ *http.Request) {
	users, err := app.Models.UserModel.GetAll()
	if err != nil {
		app.serverErrorResponse(w, err)
		return
	}

	if err := app.writeJSON(w, http.StatusOK, users, nil); err != nil {
		app.serverErrorResponse(w, err)
	}
}

func (app *Application) userDetailHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w)
		return
	}

	user, err := app.Models.UserModel.GetById(id)
	if err != nil {
		app.notFoundResponse(w)
		return
	}

	if err := app.writeJSON(w, http.StatusOK, user, nil); err != nil {
		app.serverErrorResponse(w, err)
	}
}

func (app *Application) userCreateHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}
	if err := app.readJSON(r, &input); err != nil {
		app.badRequestResponse(w, err)
		return
	}

	user := &data.User{
		Name: input.Name,
	}
	if err := app.Models.UserModel.Create(user); err != nil {
		app.badRequestResponse(w, err)
		return
	}

	if err := app.writeJSON(w, http.StatusCreated, user, nil); err != nil {
		app.serverErrorResponse(w, err)
	}
}

func (app *Application) userUpdateHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w)
		return
	}

	user, err := app.Models.UserModel.GetById(id)
	if err != nil {
		app.notFoundResponse(w)
		return
	}

	var input struct {
		Name string `json:"name"`
	}
	if err := app.readJSON(r, &input); err != nil {
		app.badRequestResponse(w, err)
		return
	}

	if input.Name != user.Name {
		user.Name = input.Name
	}
	if err := app.Models.UserModel.Update(user); err != nil {
		app.badRequestResponse(w, err)
		return
	}

	if err := app.writeJSON(w, http.StatusOK, user, nil); err != nil {
		app.serverErrorResponse(w, err)
	}
}

func (app *Application) userDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w)
		return
	}

	user, err := app.Models.UserModel.GetById(id)
	if err != nil {
		app.notFoundResponse(w)
		return
	}

	if err := app.Models.UserModel.Delete(user.ID); err != nil {
		app.serverErrorResponse(w, err)
	}

	if err := app.writeJSON(w, http.StatusNoContent, nil, nil); err != nil {
		app.serverErrorResponse(w, err)
	}
}
