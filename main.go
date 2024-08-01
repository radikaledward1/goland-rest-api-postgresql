package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/radikaledward1/golang-rest-api-postgresql/database"
	"github.com/radikaledward1/golang-rest-api-postgresql/routes"
)

func main() {
	//fmt.Println("Hello World")

	database.DbConnection()

	router := mux.NewRouter().StrictSlash(true)
	routes.RoutesRegister(router)
	log.Fatal(http.ListenAndServe(":3000", router))
}
