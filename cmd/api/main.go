package main

import (
	"github.com/tonnytg/go-apirest/internal/database"
	"github.com/tonnytg/go-apirest/internal/entity"
	"github.com/tonnytg/go-apirest/pkg/api"
)

func main() {

	entity.Personalities = []entity.Personality{
		{ID: 1, Name: "A11", History: "B11"},
		{ID: 2, Name: "A22", History: "B22"},
	}

	confDB1 := database.Config{
		Host:     "localhost",
		Port:     5432,
		User:     string("root"),
		Password: string("root"),
		DbName:   string("personalities"),
		SSLMode:  string("disable"),
	}
	db := new(database.Db)
	db.Connector(confDB1).Save(entity.Personalities)

	api.StartServer()
}
