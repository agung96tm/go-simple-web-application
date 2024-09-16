package main

import (
	"log"
	"net/http"
)

func (app *Application) serve() error {
	srv := http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	log.Printf("Starting server on port :%s", "8000")
	err := srv.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
