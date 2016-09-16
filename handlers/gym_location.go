package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
)

const GymLocationId = "gym_location_id"
const InvalidGymLocationId = "Invalid " + GymLocationId

var gym_locationFields map[string]string = map[string]string{
	"gym_location_id":    "int",
	"gym_id":             "int",
	"address_id":         "int",
	"location_name":      "string",
	"phone_number":       "string",
	"website_url":        "string",
	"in_network":         "bool",
	"monthly_member_fee": "int",
}

func GetGymLocation(w http.ResponseWriter, r *http.Request) {
	gymLocationId, message := GetId(w, r, GymLocationId)
	if message != nil {
		return
	}

	gym_location, err := datastore.GetGymLocation(gymLocationId)
	if err != nil {
		if err == sql.ErrNoRows {
			WriteJSON(w, http.StatusNotFound, APIErrorMessage{Message: "Not Found"})
		} else {
			WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: err.Error()})
		}
		return
	}

	WriteJSON(w, http.StatusOK, gym_location)
}

func GetGymLocations(w http.ResponseWriter, r *http.Request) {
	var statement string
	query := r.URL.Query()
	where, err := BuildWhere(gym_locationFields, query)
	if err != nil {
		WriteJSON(w, http.StatusNotFound, APIErrorMessage{Message: err.Error()})
		return
	}

	sort, err := BuildSort(gym_locationFields, query)
	if err != nil {
		WriteJSON(w, http.StatusNotFound, APIErrorMessage{Message: err.Error()})
		return
	}

	statement = fmt.Sprintf("%s %s", where, sort)
	statuses, err := datastore.GetGymLocationList(statement)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: "Error getting gym_location list."})
		return
	}

	WriteJSON(w, http.StatusOK, statuses)
}

func PostGymLocation(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	gym_location := &models.GymLocation{}
	err := json.Unmarshal(body, gym_location)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIErrorMessage{Message: err.Error()})
		return
	}

	created, err := datastore.CreateGymLocation(*gym_location)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: err.Error()})
		return
	}

	WriteJSON(w, http.StatusCreated, created)
}

func PutGymLocation(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	gymLocationId, err := strconv.ParseInt(mux.Vars(r)[GymLocationId], 10, 64)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIErrorMessage{Message: InvalidGymLocationId})
		return
	}

	gym_location := &models.GymLocation{}
	err = json.Unmarshal(body, gym_location)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIErrorMessage{Message: err.Error()})
		return
	}

	updated, err := datastore.UpdateGymLocation(gymLocationId, *gym_location)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, updated)
}

func DeleteGymLocation(w http.ResponseWriter, r *http.Request) {
	gymLocationId, err := strconv.ParseInt(mux.Vars(r)[GymLocationId], 10, 64)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIErrorMessage{Message: InvalidGymLocationId})
		return
	}

	err = datastore.DeleteGymLocation(gymLocationId)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, nil)
}
