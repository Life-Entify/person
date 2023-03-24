package person

import (
	"context"

	db "github.com/life-entify/person/repository/db"
	"github.com/life-entify/person/v1"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreatePerson(ctx context.Context, profile *person.Profile, nextOfKins []*person.NextOfKin, checkExistence bool) (*person.Person, error)
	DeletePersons(ctx context.Context, _ids []primitive.ObjectID) (*mongo.DeleteResult, error)
	FindPersons(ctx context.Context, keyword *person.Person, page *db.Pagination) ([]*person.Person, error)
	FindPersonById(ctx context.Context, id primitive.ObjectID) (*person.Person, error)
	FindPersonByPhone(ctx context.Context, phone string) ([]byte, error)
	FindPersonsByID(ctx context.Context, ids []primitive.ObjectID) ([]*person.Person, error)
	FindPersonsByPersonID(ctx context.Context, ids []int64) ([]*person.Person, error)
	FindPersonProfile(ctx context.Context, _id string) ([]byte, error)
	UpdatePerson(ctx context.Context, _id primitive.ObjectID, profile *person.Profile) (*mongo.UpdateResult, error)
	SetNextOfKins(ctx context.Context, personID int64, nextOfKins []*person.NextOfKin) (*mongo.UpdateResult, error)
}
