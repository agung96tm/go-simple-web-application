package main

import "github.com/graphql-go/graphql"

var userResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserResponse",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var userListResponse = graphql.NewList(userResponse)

var userParam = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

var userCreateRequest = graphql.FieldConfigArgument{
	"name": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

var userUpdateRequest = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"name": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
