package main

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
)

func (app *Application) readParams(params graphql.ResolveParams, v interface{}) error {
	jsonData, err := json.Marshal(params.Args)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonData, v); err != nil {
		return err
	}
	return nil
}
