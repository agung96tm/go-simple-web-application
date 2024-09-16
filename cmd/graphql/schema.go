package main

import (
	"github.com/agung96tm/go-simple-web-application/internal/graph"
)

func (app *Application) routes() graph.Fields {
	fields := graph.NewFields()

	fields.HandleQuery("listUsers", nil, app.listUsersHandler, userListResponse)
	fields.HandleMutation("createUser", userCreateRequest, app.createUserHandler, userResponse)
	fields.HandleQuery("detailUser", userParam, app.detailUserHandler, userResponse)
	fields.HandleMutation("updateUser", userUpdateRequest, app.updateUserHandler, userResponse)
	fields.HandleMutation("deleteUser", userParam, app.deleteUserHandler, userResponse)

	return fields
}
