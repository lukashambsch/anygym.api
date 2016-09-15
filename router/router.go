package router

import (
	"fmt"
	"net/http"
	"os"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/lukashambsch/gym-all-over/handlers"
)

var V1URLBase string = "/api/v1"

func Load() http.Handler {

	r := mux.NewRouter().StrictSlash(true)

	// Status endpoints
	s := r.PathPrefix(fmt.Sprintf("%s%s", V1URLBase, "/statuses")).Subrouter()

	s.HandleFunc("/", handlers.GetStatuses).
		Methods("GET")
	s.HandleFunc("/{status_id}/", handlers.GetStatus).
		Methods("GET")
	s.HandleFunc("/", handlers.PostStatus).
		Methods("POST")
	s.HandleFunc("/{status_id}/", handlers.PutStatus).
		Methods("PUT")
	s.HandleFunc("/{status_id}/", handlers.DeleteStatus).
		Methods("DELETE")

	// Visit endpoints
	v := r.PathPrefix(fmt.Sprintf("%s%s", V1URLBase, "/visits")).Subrouter()

	v.HandleFunc("/", handlers.GetVisits).
		Methods("GET")
	v.HandleFunc("/{visit_id}/", handlers.GetVisit).
		Methods("GET")
	v.HandleFunc("/", handlers.PostVisit).
		Methods("POST")
	v.HandleFunc("/{visit_id}/", handlers.PutVisit).
		Methods("PUT")
	v.HandleFunc("/{visit_id}/", handlers.DeleteVisit).
		Methods("DELETE")

	router := ghandlers.LoggingHandler(os.Stdout, r)
	router = ghandlers.CORS()(router)

	return router
}
