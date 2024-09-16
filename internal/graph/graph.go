package graph

import "github.com/graphql-go/graphql"

type Fields struct {
	queryFields    graphql.Fields
	mutationFields graphql.Fields
}

func NewFields() Fields {
	return Fields{
		queryFields:    graphql.Fields{},
		mutationFields: graphql.Fields{},
	}
}

func (f *Fields) HandleQuery(path string, args graphql.FieldConfigArgument, resolve func(p graphql.ResolveParams) (interface{}, error), out graphql.Output) {
	f.queryFields[path] = &graphql.Field{
		Type:    out,
		Args:    args,
		Resolve: resolve,
	}
}

func (f *Fields) GenerateQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: f.queryFields})
}

func (f *Fields) HandleMutation(path string, args graphql.FieldConfigArgument, resolve func(p graphql.ResolveParams) (interface{}, error), out graphql.Output) {
	f.mutationFields[path] = &graphql.Field{
		Type:    out,
		Args:    args,
		Resolve: resolve,
	}
}

func (f *Fields) GenerateMutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{Name: "Mutation", Fields: f.mutationFields})
}
