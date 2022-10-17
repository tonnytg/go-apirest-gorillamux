package main

import (
	"github.com/tonnytg/go-apirest/models"
	"github.com/tonnytg/go-apirest/routes"
)

func main() {

	models.Personaliades = []models.Personalidade{
		{ID: 1, Nome: "A1", Historia: "B1"},
		{ID: 2, Nome: "A2", Historia: "B2"},
	}

	routes.HandleRequest()
}
