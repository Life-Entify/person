package person

import (
	"github.com/graphql-go/graphql"
)

var KeywordPersonInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "KeywordPersonInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"person_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"profile": &graphql.InputObjectFieldConfig{
			Type: KeywordProfileInputType,
		},
	},
})
var PersonInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PersonInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"person_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"profile": &graphql.InputObjectFieldConfig{
			Type: ProfileInputType,
		},
	},
})
var AddressInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "AddressInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"street": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"town": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"lga": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"nstate": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"country": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var NextOfKinArrayInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "NextOfKinArrayInputType",
	Fields: graphql.InputObjectField{
		Type: graphql.NewList(NextOfKinInputType),
	},
})
var NextOfKinInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "NextOfKinInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"relationship": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"next_of_kin": &graphql.InputObjectFieldConfig{
			Type: ProfileInputType,
		},
	},
})
var KeywordProfileInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "KeywordProfileInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"last_name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"first_name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"middle_name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"phone_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"national_identity": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
var ProfileInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ProfileInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"last_name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"first_name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"middle_name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"phone_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"occupation": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"dob": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"gender": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"national_identity": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"addresses": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(AddressInputType),
		},
	},
})
