package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Go REST API !!!")
}

func main() {
	//fmt.Println("Hello World")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Welcome)

	log.Fatal(http.ListenAndServe(":3000", router))
}
