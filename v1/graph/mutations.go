package person

import (
	"github.com/graphql-go/graphql"
	schemas "github.com/life-entify/person/v1/graph/schemas"
)

func CreatePerson(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Create Person",
		Type:        schemas.PersonType,
		Args: graphql.FieldConfigArgument{
			"profile": &graphql.ArgumentConfig{
				Type: schemas.ProfileInputType,
			},
		},
		Resolve: resolver,
	}
}
func UpdatePerson(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "UpdatePerson Person",
		Type:        schemas.PersonType,
		Args: graphql.FieldConfigArgument{
			"_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"profile": &graphql.ArgumentConfig{
				Type: schemas.ProfileInputType,
			},
		},
		Resolve: resolver,
	}
}
