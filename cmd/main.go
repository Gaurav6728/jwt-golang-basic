package main

import (
	"log"
	"net/http"

	"github.com/gaurav/golang-jwt-project/usecase"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/login", usecase.Login).Methods(http.MethodPost)
	router.HandleFunc("/protected-route", usecase.Protected).Methods(http.MethodGet)

	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal("error starting the server")
		return
	}
	log.Print("server started")

}
