package main

import (
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

func (app *Application) serve() error {
	s := grpc.NewServer(grpc.KeepaliveParams(
		keepalive.ServerParameters{
			MaxConnectionAge:  5 * time.Minute,
			Timeout:           15 * time.Second,
			MaxConnectionIdle: 5 * time.Minute,
			Time:              15 * time.Minute,
		},
	))

	app.Grpc = s
	app.routes()

	// listen
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", app.Config.Port))
	if err != nil {
		return errors.New("failed to listen: " + err.Error())
	}

	app.Logger.Printf("Starting RPC server at %s\n", app.Config.Port)
	err = s.Serve(listen)
	if err != nil {
		return errors.New("failed to serve: " + err.Error())
	}

	return err
}
