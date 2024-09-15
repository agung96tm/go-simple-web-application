package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *Application) serve() error {
	srv := http.Server{
		Addr:         fmt.Sprintf(":%s", app.Config.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	app.Logger.Printf("Starting server on port %s", app.Config.Port)
	err := srv.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
