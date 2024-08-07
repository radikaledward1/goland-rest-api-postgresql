package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/radikaledward1/golang-rest-api-postgresql/database"
	"github.com/radikaledward1/golang-rest-api-postgresql/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	database.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)
}

func AddUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	createdUser := database.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(&user)

}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete Users Handler")
}
