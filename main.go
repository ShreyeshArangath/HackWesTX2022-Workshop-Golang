package hackwestx_workshop

import "github.com/gorilla/mux"

func GetUsers() {

}

func GetUserById() {
	
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/users", GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", GetUserById).Methods("GET")
}
