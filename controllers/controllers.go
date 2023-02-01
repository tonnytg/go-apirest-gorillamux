package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/tonnytg/go-apirest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

// Create Token with Json
func Token(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body := r.Body
	var user models.User
	json.NewDecoder(body).Decode(&user)
	token := models.CreateToken(user.Username, user.Password)
	json.NewEncoder(w).Encode(token)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personalidade := range models.Personaliades {
		if strconv.Itoa(personalidade.ID) == id {
			json.NewEncoder(w).Encode(personalidade)
		}
	}
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Criando")
}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deletando")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personaliades)
}
