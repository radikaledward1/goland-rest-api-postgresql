package routes

import (
	"github.com/gorilla/mux"
	"github.com/radikaledward1/golang-rest-api-postgresql/services"
)

func RoutesRegister(router *mux.Router) {
	router.HandleFunc("/", services.Welcome).Methods("GET")
	router.HandleFunc("/users", services.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/add", services.AddUsersHandler).Methods("POST")
	router.HandleFunc("/users/delete/{id}", services.DeleteUsersHandler).Methods("POST")
}
