package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ShreyeshArangath/hackwestx-workshop/internal/user/entity"
	"github.com/ShreyeshArangath/hackwestx-workshop/internal/user/repository"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type UserServiceController interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	repo   repository.Repository
	Logger log.Logger
}

func NewController(repo repository.Repository, logger log.Logger) UserServiceController {
	return controller{repo, logger}
}

func (c controller) GetUsers(w http.ResponseWriter, r *http.Request) {
	//TODO: needs to be implemented
}

func (c controller) GetUserById(w http.ResponseWriter, r *http.Request) {
	//TODO: needs to be implemented
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user entity.User
	id := params["id"]
	fmt.Println(id, ": UserID")
	json.NewEncoder(w).Encode(user)
	user, err := c.repo.Get(context.TODO(), id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(user)
}

func (c controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	//TODO: needs to be implemented
	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)
	c.repo.Create(context.TODO(), user)
	json.NewEncoder(w).Encode(user)
}

func (c controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	//TODO: needs to be implemented
}

func (c controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	//TODO: needs to be implemented
}
