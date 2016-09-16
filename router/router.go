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

	// Member endpoints
	m := r.PathPrefix(fmt.Sprintf("%s%s", V1URLBase, "/members")).Subrouter()

	m.HandleFunc("/", handlers.GetMembers).
		Methods("GET")
	m.HandleFunc("/{member_id}/", handlers.GetMember).
		Methods("GET")
	m.HandleFunc("/", handlers.PostMember).
		Methods("POST")
	m.HandleFunc("/{member_id}/", handlers.PutMember).
		Methods("PUT")
	m.HandleFunc("/{member_id}/", handlers.DeleteMember).
		Methods("DELETE")

	// GymLocation endpoints
	gl := r.PathPrefix(fmt.Sprintf("%s%s", V1URLBase, "/gym_locations")).Subrouter()

	gl.HandleFunc("/", handlers.GetGymLocations).
		Methods("GET")
	gl.HandleFunc("/{gym_location_id}/", handlers.GetGymLocation).
		Methods("GET")
	gl.HandleFunc("/", handlers.PostGymLocation).
		Methods("POST")
	gl.HandleFunc("/{gym_location_id}/", handlers.PutGymLocation).
		Methods("PUT")
	gl.HandleFunc("/{gym_location_id}/", handlers.DeleteGymLocation).
		Methods("DELETE")

	router := ghandlers.LoggingHandler(os.Stdout, r)
	router = ghandlers.CORS()(router)

	return router
}
