package user

import (
	"github.com/ShreyeshArangath/hackwestx-workshop/internal/user/controller"
	"github.com/ShreyeshArangath/hackwestx-workshop/internal/user/repository"
	"github.com/ShreyeshArangath/hackwestx-workshop/pkg"
	"github.com/ShreyeshArangath/hackwestx-workshop/pkg/db"
	"github.com/gorilla/mux"
	"log"
)

type userservice struct {
	db     db.DBhelper
	logger log.Logger
}

func NewUserService(db db.DBhelper, logger log.Logger) pkg.Service {
	return userservice{
		db:     db,
		logger: logger,
	}
}

func (u userservice) InitializeService(r *mux.Router) {
	// Dependency injection
	repository := repository.NewRepository(u.db, u.logger)
	controller := controller.NewController(repository, u.logger)

	// Register our routes
	r.HandleFunc("/api/users", controller.GetUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", controller.GetUserById).Methods("GET")
	r.HandleFunc("/api/users", controller.CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", controller.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", controller.DeleteUser).Methods("DELETE")

}
