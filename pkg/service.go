package pkg

import "github.com/gorilla/mux"

type Service interface {
	InitializeService(r *mux.Router)
}
