package routes

import (
	"github.com/gorilla/mux"
	"github.com/tonnytg/go-apirest/controllers"
)

func LoadHandlers(r *mux.Router) {

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/token", controllers.Token)
	r.HandleFunc("/auth", controllers.Auth)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("GET")
	r.HandleFunc("/api/personality", controllers.CreatePersonality).Methods("POST")
	r.HandleFunc("/api/personality/{id}", controllers.GetPersonality).Methods("GET")
	r.HandleFunc("/api/personality/{id}", controllers.DeletePersonality).Methods("DELETE")
}
