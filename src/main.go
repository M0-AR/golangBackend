package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/B1gDaddyKane/golangBackend/src/controllers"
)

func homePage(w http.ResponseWriter, r *http.Request) {

}

func signUp(w http.ResponseWriter, r *http.Request) {

}

func handleRequest() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/signUp", signUp)

	// Dashboard -> Service
	router.HandleFunc("/dashboard/services", controllers.GetServices).Methods("GET")
	router.HandleFunc("/dashboard/services/{id}", controllers.GetServiceById).Methods("GET")

	fmt.Print("Server at PORT 10000")
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	handleRequest()
}
