package person

import (
	"github.com/graphql-go/graphql"
)

var AddressType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Address",
	Fields: graphql.Fields{
		"street": &graphql.Field{
			Type: graphql.String,
		},
		"town": &graphql.Field{
			Type: graphql.String,
		},
		"lga": &graphql.Field{
			Type: graphql.String,
		},
		"nstate": &graphql.Field{
			Type: graphql.String,
		},
		"country": &graphql.Field{
			Type: graphql.String,
		},
		"_id": &graphql.Field{
			Type: graphql.String,
		},
	},
})
var ProfileType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Profile",
	Fields: graphql.Fields{
		"last_name": &graphql.Field{
			Type: graphql.String,
		},
		"first_name": &graphql.Field{
			Type: graphql.String,
		},
		"middle_name": &graphql.Field{
			Type: graphql.String,
		},
		"phone_number": &graphql.Field{
			Type: graphql.String,
		},
		"occupation": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"gender": &graphql.Field{
			Type: graphql.String,
		},
		"dob": &graphql.Field{
			Type: graphql.String,
		},
		"national_identity": &graphql.Field{
			Type: graphql.String,
		},
		"addresses": &graphql.Field{
			Type: graphql.NewList(AddressType),
		},
	},
})
var PersonType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PersonType",
	Fields: graphql.Fields{
		"person_id": &graphql.Field{
			Type: graphql.Int,
		},
		"_id": &graphql.Field{
			Type: graphql.String,
		},
		"profile": &graphql.Field{
			Type: ProfileType,
		},
		"next_of_kins": &graphql.Field{
			Type: graphql.NewList(NextOfKinType),
		},
	},
})

var NextOfKinType = graphql.NewObject(graphql.ObjectConfig{
	Name: "NextOfKin",
	Fields: graphql.Fields{
		"person_id": &graphql.Field{
			Type: graphql.Int,
		},
		"relationship": &graphql.Field{
			Type: graphql.String,
		},
	},
})
