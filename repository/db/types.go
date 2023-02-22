package person

import "github.com/life-entify/person/v1"

type Pagination struct {
	Limit, Skip int64
}

type Keyword struct {
	Profile person.Person   `json:"profile"`
}
