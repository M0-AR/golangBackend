package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"

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

	fmt.Print("Server start at PORT 10000\n")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":10000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func main() {
	handleRequest()
}
