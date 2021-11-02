package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/B1gDaddyKane/golangBackend/database"
	"github.com/B1gDaddyKane/golangBackend/middleware"
	"github.com/B1gDaddyKane/golangBackend/models"
	"github.com/gorilla/mux"
)

var db = database.NewMongoDB("mongodb+srv://admin:admin@bookus0.4yn8c.mongodb.net/BookUs0?retryWrites=true&w=majority")

func homePage(res http.ResponseWriter, req *http.Request) {

}

// signs an user up and creates a new user in the database
/* TODO:
1. Create an actual data base
2. Move database connection and functionalities to a seperate file
3. A login function and a JWT token function
*/
func signUp(res http.ResponseWriter, req *http.Request) {
	// sets the response header
	res.Header().Set("Content-Type", "application/json")
	var user models.User
	// decodes the request body into our user struct
	json.NewDecoder(req.Body).Decode(&user)
	result, err := db.CreateUser(database.UserDB{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  middleware.GetHash([]byte(user.Password)),
		Created:   time.Now().Unix(),
	})
	if err != nil {
		return
	}
	json.NewEncoder(res).Encode(result)
}

func handleRequest() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/signUp", signUp).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	handleRequest()
}
