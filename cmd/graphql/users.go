package main

import (
	"github.com/agung96tm/go-simple-web-application/internal/data"
	"github.com/graphql-go/graphql"
)

func (app *Application) listUsersHandler(_ graphql.ResolveParams) (interface{}, error) {
	users, err := app.Models.UserModel.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (app *Application) detailUserHandler(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)

	user, err := app.Models.UserModel.GetById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (app *Application) createUserHandler(params graphql.ResolveParams) (interface{}, error) {
	var input struct {
		Name string `json:"name"`
	}
	err := app.readParams(params, &input)
	if err != nil {
		return nil, err
	}

	user := &data.User{Name: input.Name}
	err = app.Models.UserModel.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (app *Application) updateUserHandler(params graphql.ResolveParams) (interface{}, error) {
	var input struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	err := app.readParams(params, &input)
	if err != nil {
		return nil, err
	}

	user, err := app.Models.UserModel.GetById(input.ID)
	if err != nil {
		return nil, err
	}
	if user.Name != input.Name {
		user.Name = input.Name
	}

	err = app.Models.UserModel.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (app *Application) deleteUserHandler(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)

	user, err := app.Models.UserModel.GetById(id)
	if err != nil {
		return nil, err
	}

	err = app.Models.UserModel.Delete(user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
