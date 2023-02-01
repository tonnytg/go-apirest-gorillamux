package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/tonnytg/go-apirest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("ENV_VERSION")
	// get time UTC from server
	time := time.Now().UTC()
	path := "http://localhost:8081"
	log.Printf("[%s] - %s %d %s", r.RemoteAddr, r.Method, 200, path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"version": version, "time": time.String()})
}

func Token(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body := r.Body
	var user models.User
	json.NewDecoder(body).Decode(&user)
	if user.Username == "" || user.Password == "" {
		log.Printf("[%s] %s - Username or password is empty", r.RemoteAddr, r.Method)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Username or password is empty\n")
		return
	}
	token := models.CreateToken(user.Username, user.Password)
	log.Printf("[%s] %s - Token created", r.RemoteAddr, r.Method)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

func Auth(w http.ResponseWriter, r *http.Request) {
	// extract Authorization: Bearer from request

	token := r.Header.Get("Authorization")
	log.Println(token)
	fmt.Println("Auth")
}

func GetPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personality := range models.Personalities {
		if strconv.Itoa(personality.ID) == id {
			json.NewEncoder(w).Encode(personality)
		}
	}
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	// TODO: Create a personality
	fmt.Println("Creating")
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	// TODO: Delete Personalities
	fmt.Println("Deleting")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}
