package main

import (
	"github.com/ShreyeshArangath/hackwestx-workshop/internal/user/entity"
	"github.com/gorilla/mux"
	"net/http"
)

// Main HTTP Router
func main() {

	// Read in the environment variables - MAKE SURE TO NOT LET YOUR .ENV GET INTO GIT

	// Initiate a new router
	router := mux.NewRouter()

	// Handling the routes

	userController := controller.New
	// User Service
	router.HandleFunc("/api/users").Methods("GET")
	router.HandleFunc("/api/users/{id}", entity.GetUserById).Methods("GET")

	// To start our HTTP server
	http.ListenAndServe(":8080", router)
}
