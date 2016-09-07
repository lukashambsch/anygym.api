package router

import (
	"fmt"

	"github.com/gorilla/mux"

	"github.com/lukashambsch/gym-all-over/handlers"
)

var V1URLBase string = "/api/v1"

func Load() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)

	// Status endpoints
	s := r.PathPrefix(fmt.Sprintf("%s%s", V1URLBase, "/statuses")).Subrouter()

	s.HandleFunc("/", handlers.GetStatuses).
		Methods("GET")
	s.HandleFunc("/{status_id}", handlers.GetStatus).
		Methods("GET")
	s.HandleFunc("/", handlers.PostStatus).
		Methods("POST")
	s.HandleFunc("/{status_id}", handlers.PutStatus).
		Methods("PUT")
	s.HandleFunc("/{status_id}", handlers.DeleteStatus).
		Methods("DELETE")

	return r
}
