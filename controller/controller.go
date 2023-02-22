package person

import (
	config "github.com/life-entify/person/config"
	repo "github.com/life-entify/person/repository"
	db "github.com/life-entify/person/repository/db"
)

type Controller struct {
	repo.Repository
	Config *config.Config
}

const (
	Mongo    = "MONGODB"
	MySQL    = "MYSQL"
	PostGres = "POSTGRES"
)

func NewController(config *config.Config) *Controller {
	return &Controller{
		db.NewMongoDB(config),
		config,
	}
}
