package main

import (
	"fmt"
	"github.com/B1gDaddyKane/golangBackend/lib/config"
	"github.com/B1gDaddyKane/golangBackend/lib/db"
	"github.com/B1gDaddyKane/golangBackend/lib/logging"
	"github.com/B1gDaddyKane/golangBackend/src/dao"
	httpDelivery "github.com/B1gDaddyKane/golangBackend/src/delivery/http"
	"github.com/B1gDaddyKane/golangBackend/src/model"
	"github.com/B1gDaddyKane/golangBackend/src/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	"log"

	//"github.com/gorilla/handlers"
	//"github.com/gorilla/mux"
	"github.com/labstack/echo/v4/middleware"
)

//func homePage(w http.ResponseWriter, r *http.Request) {
//
//}
//
//func signUp(w http.ResponseWriter, r *http.Request) {
//
//}
//
//func handleRequest() {
//
//	router := mux.NewRouter().StrictSlash(true)
//
//	router.HandleFunc("/", homePage)
//	router.HandleFunc("/signUp", signUp)
//
//	// Dashboard -> Service
//	router.HandleFunc("/dashboard/usecase", controllers.GetServices).Methods("GET")
//	router.HandleFunc("/dashboard/usecase/{id}", controllers.GetServiceById).Methods("GET")
//
//	fmt.Print("Server start at PORT 10000\n")
//
//	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
//	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
//	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
//
//	log.Fatal(http.ListenAndServe(":10000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
//}

// To initialize viper config
func init() {
	config.SetConfigFile("config", "lib/config", "json")
}

func main() {
	//handleRequest()

	envConfig := getConfig()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.Use(logging.MiddlewareLogging)

	// Mongo
	mongo, err := db.Connect(envConfig.Mongo)
	if err != nil {
		log.Println(err)
		return
	}

	serviceRepo := dao.NewServiceDAO(mongo)
	serviceUseCase := usecase.NewServiceUseCase(&envConfig, serviceRepo)

	// Router
	httpDelivery.NewRouter(e, serviceUseCase)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s%s%v", envConfig.Host, ":", envConfig.Port)))
}

func getConfig() model.EnvConfig {

	return model.EnvConfig{
		Host: config.GetString("host.address"),
		Port: config.GetInt("host.port"),
		Mongo: db.MongoConfig{
			Timeout:  config.GetInt("database.mongodb.timeout"),
			DBName:   config.GetString("database.mongodb.dbname"),
			Username: config.GetString("database.mongodb.user"),
			Password: config.GetString("database.mongodb.password"),
			Host:     config.GetString("database.mongodb.host"),
			Port:     config.GetString("database.mongodb.port"),
		},
	}
}
