package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/radikaledward1/golang-rest-api-postgresql/database"
	"github.com/radikaledward1/golang-rest-api-postgresql/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	database.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	/*userId, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid Id")
		return
	}

	var Users []models.User
	database.DB.Find(&Users)

	for _, user := range Users {
		if user.ID == uint(userId) {
			json.NewEncoder(w).Encode(user)
		}
	} */

	var user models.User
	database.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User not found")
		return
	}

	json.NewEncoder(w).Encode(user)
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
	params := mux.Vars(r)

	var user models.User
	database.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User not found")
		return
	}

	database.DB.Delete(&user) //virtual deleted, deletedAt field updated
	//database.DB.Unscoped.Delete(user) //removed forever, WARNING
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User deleted successfuly")
}
