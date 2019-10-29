package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	baseChromePath = "/home/chrome/Downloads"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/statement", getStatemenHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))

}
