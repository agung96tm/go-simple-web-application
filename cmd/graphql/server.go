package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"net/http"
)

func (app *Application) serve() error {
	fields := app.routes()
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    fields.GenerateQuery(),
		Mutation: fields.GenerateMutation(),
	})
	if err != nil {
		return err
	}

	app.Logger.Printf("Starting server on port %s", app.Config.Port)
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/graphql", h)
	return http.ListenAndServe(fmt.Sprintf(":%s", app.Config.Port), nil)
}
