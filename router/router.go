package router

import (
	"fmt"
	"net/http"
	"os"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/lukashambsch/anygym.api/handlers"
)

const V1URLBase string = "/api/v1"

func Load() http.Handler {

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc(fmt.Sprintf("%s%s", V1URLBase, "/authenticate"), handlers.Login).Methods("POST")
	r.HandleFunc(fmt.Sprintf("%s%s", V1URLBase, "/logout"), handlers.Logout)

	// Status endpoints
	statuses := fmt.Sprintf("%s/statuses", V1URLBase)

	r.HandleFunc(statuses, handlers.GetStatuses).
		Methods("GET")
	r.HandleFunc(fmt.Sprintf("%s/{status_id}", statuses), handlers.GetStatus).
		Methods("GET")
	r.HandleFunc(statuses, handlers.PostStatus).
		Methods("POST")
	r.HandleFunc(fmt.Sprintf("%s/{status_id}", statuses), handlers.PutStatus).
		Methods("PUT")
	r.HandleFunc(fmt.Sprintf("%s/{status_id}", statuses), handlers.DeleteStatus).
		Methods("DELETE")

	// Visit endpoints
	visits := fmt.Sprintf("%s/visits", V1URLBase)

	r.HandleFunc(visits, handlers.GetVisits).
		Methods("GET")
	r.HandleFunc(fmt.Sprintf("%s/{visit_id}", visits), handlers.GetVisit).
		Methods("GET")
	r.HandleFunc(visits, handlers.PostVisit).
		Methods("POST")
	r.HandleFunc(fmt.Sprintf("%s/{visit_id}", visits), handlers.PutVisit).
		Methods("PUT")
	r.HandleFunc(fmt.Sprintf("%s/{visit_id}", visits), handlers.DeleteVisit).
		Methods("DELETE")

	// Member endpoints
	members := fmt.Sprintf("%s/members", V1URLBase)

	r.HandleFunc(members, handlers.GetMembers).
		Methods("GET")
	r.HandleFunc(fmt.Sprintf("%s/{member_id}", members), handlers.GetMember).
		Methods("GET")
	r.HandleFunc(members, handlers.PostMember).
		Methods("POST")
	r.HandleFunc(fmt.Sprintf("%s/{member_id}", members), handlers.PutMember).
		Methods("PUT")
	r.HandleFunc(fmt.Sprintf("%s/{member_id}", members), handlers.DeleteMember).
		Methods("DELETE")

	// GymLocation endpoints
	gymLocations := fmt.Sprintf("%s/gym_locations", V1URLBase)

	r.HandleFunc(gymLocations, handlers.GetGymLocations).
		Methods("GET")
	r.HandleFunc(fmt.Sprintf("%s/{gym_location_id}", gymLocations), handlers.GetGymLocation).
		Methods("GET")
	r.HandleFunc(gymLocations, handlers.PostGymLocation).
		Methods("POST")
	r.HandleFunc(fmt.Sprintf("%s/{gym_location_id}", gymLocations), handlers.PutGymLocation).
		Methods("PUT")
	r.HandleFunc(fmt.Sprintf("%s/{gym_location_id}", gymLocations), handlers.DeleteGymLocation).
		Methods("DELETE")

	// User endpoints
	users := fmt.Sprintf("%s/users", V1URLBase)

	r.HandleFunc(users, handlers.GetUsers).
		Methods("GET")
	r.HandleFunc(fmt.Sprintf("%s/{user_id}", users), handlers.GetUser).
		Methods("GET")
	r.HandleFunc(users, handlers.PostUser).
		Methods("POST")
	r.HandleFunc(fmt.Sprintf("%s/{user_id}", users), handlers.PutUser).
		Methods("PUT")
	r.HandleFunc(fmt.Sprintf("%s/{user_id}", users), handlers.DeleteUser).
		Methods("DELETE")

	router := ghandlers.LoggingHandler(os.Stdout, r)
	router = handlers.CORS(router)
	router = handlers.VerifyToken(router)

	return router
}
