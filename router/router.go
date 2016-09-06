package router

import (
	//"fmt"

	//"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"

	"github.com/lukashambsch/gym-all-over/handlers"
)

var V1URLBase string = "/api/v1"

func Load() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api/v1/statuses", handlers.GetStatuses).
		Methods("GET")
	r.HandleFunc("/api/v1/statuses/{status_id}", handlers.GetStatus).
		Methods("GET")
	r.HandleFunc("/api/v1/statuses", handlers.PostStatus).
		Methods("POST")
	r.HandleFunc("/api/v1/statuses/{status_id}", handlers.PutStatus).
		Methods("PUT")
	r.HandleFunc("/api/v1/statuses/{status_id}", handlers.DeleteStatus).
		Methods("DELETE")

	return r
}
