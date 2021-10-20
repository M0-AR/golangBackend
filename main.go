package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {

}

func signUp(w http.ResponseWriter, r *http.Request) {

}

func handleRequest() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/signUp", signUp)

	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	handleRequest()
}
