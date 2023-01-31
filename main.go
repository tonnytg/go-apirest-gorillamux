package main

import (
	"github.com/tonnytg/go-apirest/database"
	"github.com/tonnytg/go-apirest/models"
	"github.com/tonnytg/go-apirest/routes"
)

func main() {

	models.Personaliades = []models.Personalidade{
		{ID: 3, Nome: "A1", Historia: "B1"},
		{ID: 4, Nome: "A2", Historia: "B2"},
	}

	db := new(database.Db)

	routes.HandleRequest()
}
