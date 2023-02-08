package main

import (
	"github.com/gorilla/mux"
	"github.com/tonnytg/go-apirest/internal/database"
	"github.com/tonnytg/go-apirest/pkg/api/models"
	"github.com/tonnytg/go-apirest/pkg/api/routes"
	"log"
	"net/http"
	"os"
)

func main() {

	models.Personalities = []models.Personality{
		{ID: 1, Name: "A11", History: "B11"},
		{ID: 2, Name: "A22", History: "B22"},
	}

	db := new(database.Db)
	db.Connector().Save(models.Personalities)

	PORT := os.Getenv("ENV_PORT")
	if PORT == "" {
		PORT = "8080"
	}

	r := mux.NewRouter()
	routes.LoadHandlers(r)
	log.Fatal(http.ListenAndServe(":8081", r))
}
