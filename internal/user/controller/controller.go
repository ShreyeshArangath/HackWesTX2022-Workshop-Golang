package controller

import (
	"github.com/ShreyeshArangath/hackwestx-workshop/internal/user/entity"
	"github.com/ShreyeshArangath/hackwestx-workshop/internal/user/repository"
	"log"
	"net/http"
)

type UserServiceController interface {
	GetUsers(w http.ResponseWriter, r *http.Request) (entity.User, error)
	GetUserById(w http.ResponseWriter, r *http.Request) (entity.User, error)
	CreateUser(w http.ResponseWriter, r *http.Request) (entity.User, error)
	UpdateUser(w http.ResponseWriter, r *http.Request) (entity.User, error)
	DeleteUser(w http.ResponseWriter, r *http.Request) (entity.User, error)
}

type controller struct {
	repo   repository.Repository
	Logger log.Logger
}

func NewController(repo repository.Repository, logger log.Logger) UserServiceController {
	return controller{repo, logger}
}

func (c controller) GetUsers(w http.ResponseWriter, r *http.Request) (entity.User, error) {
	return entity.User{}, nil
}

func (c controller) GetUserById(w http.ResponseWriter, r *http.Request) (entity.User, error) {
	return entity.User{}, nil
}

func (c controller) CreateUser(w http.ResponseWriter, r *http.Request) (entity.User, error) {
	return entity.User{}, nil
}

func (c controller) UpdateUser(w http.ResponseWriter, r *http.Request) (entity.User, error) {
	return entity.User{}, nil
}

func (c controller) DeleteUser(w http.ResponseWriter, r *http.Request) (entity.User, error) {
	return entity.User{}, nil
}
