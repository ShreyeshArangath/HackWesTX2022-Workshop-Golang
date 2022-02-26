package main

import (
	"github.com/ShreyeshArangath/hackwestx-workshop/internal/user"
	db "github.com/ShreyeshArangath/hackwestx-workshop/pkg/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

// Main HTTP Router
func main() {

	// Read in the environment variables - MAKE SURE TO NOT LET YOUR .ENV GET INTO GIT
	logger := log.Logger{}
	password := os.Getenv("TTUDSC_PASSWORD")
	mongoURI := "mongodb+srv://ttudsc-admin:" + password + "@cluster0.0n0rh.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	mongoDBName := "hackwestx-workshop"
	database := db.NewDatabase(mongoURI, mongoDBName, logger)
	// Initiate a new router
	router := mux.NewRouter()

	// Get the microservices
	userService := user.NewUserService(database, logger)
	userService.InitializeService(router)

	// To start our HTTP server
	http.ListenAndServe(":8080", router)
}
