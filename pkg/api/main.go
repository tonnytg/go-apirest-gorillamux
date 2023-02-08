package api

import (
	"github.com/gorilla/mux"
	"github.com/tonnytg/go-apirest/pkg/api/routes"
	"log"
	"net/http"
	"os"
)

func StartServer() {

	PORT := os.Getenv("ENV_PORT")
	if PORT == "" {
		PORT = "8080"
	}

	r := mux.NewRouter()
	routes.LoadHandlers(r)
	log.Fatal(http.ListenAndServe(":8081", r))
}
