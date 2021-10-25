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
	// Todo: Can't not differentiate between these to GET method
	//router.HandleFunc("/dashboard/services", controllers.GetServices).Methods("GET")
	router.HandleFunc("/dashboard/services", controllers.GetServiceById).Methods("GET") // Test done by postman

	fmt.Print("Server start at PORT 10000\n")
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	handleRequest()
}
