package main

import (
	"github.com/agung96tm/go-simple-web-application/internal/proto"
)

func (app *Application) routes() {
	proto.RegisterUserServiceServer(app.Grpc, NewUserService(app))
}
