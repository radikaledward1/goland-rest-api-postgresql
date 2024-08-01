package routes

import (
	"github.com/gorilla/mux"
	"github.com/radikaledward1/golang-rest-api-postgresql/services"
)

func RoutesRegister(router *mux.Router) {
	router.HandleFunc("/", services.Welcome)
}
