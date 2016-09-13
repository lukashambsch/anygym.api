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

const VisitId = "visit_id"
const InvalidVisitId = "Invalid " + VisitId

var visitFields map[string]string = map[string]string{
	"visit_id":        "int",
	"member_id":       "string",
	"gym_location_id": "string",
	"status_id":       "string",
	"created_on":      "date",
	"modified_on":     "date",
}

func GetVisit(w http.ResponseWriter, r *http.Request) {
	statusId, message := GetId(w, r, VisitId)
	if message != nil {
		return
	}

	status, err := datastore.GetVisit(statusId)
	if err != nil {
		if err == sql.ErrNoRows {
			WriteJSON(w, http.StatusNotFound, APIErrorMessage{Message: "Not Found"})
		} else {
			WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: err.Error()})
		}
		return
	}

	WriteJSON(w, http.StatusOK, status)
}

func GetVisits(w http.ResponseWriter, r *http.Request) {
	var statement string
	query := r.URL.Query()
	where, err := BuildWhere(visitFields, query)
	if err != nil {
		WriteJSON(w, http.StatusNotFound, APIErrorMessage{Message: err.Error()})
		return
	}

	sort, err := BuildSort(visitFields, query)
	if err != nil {
		WriteJSON(w, http.StatusNotFound, APIErrorMessage{Message: err.Error()})
		return
	}

	statement = fmt.Sprintf("%s %s", where, sort)
	statuses, err := datastore.GetVisitList(statement)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: "Error getting status list."})
		return
	}

	WriteJSON(w, http.StatusOK, statuses)
}

func PostVisit(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	status := &models.Visit{}
	err := json.Unmarshal(body, status)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIErrorMessage{Message: err.Error()})
		return
	}

	created, err := datastore.CreateVisit(*status)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: err.Error()})
		return
	}

	WriteJSON(w, http.StatusCreated, created)
}

func PutVisit(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	statusId, err := strconv.ParseInt(mux.Vars(r)[VisitId], 10, 64)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIErrorMessage{Message: InvalidVisitId})
		return
	}

	status := &models.Visit{}
	err = json.Unmarshal(body, status)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIErrorMessage{Message: err.Error()})
		return
	}

	updated, err := datastore.UpdateVisit(statusId, *status)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, updated)
}

func DeleteVisit(w http.ResponseWriter, r *http.Request) {
	statusId, err := strconv.ParseInt(mux.Vars(r)[VisitId], 10, 64)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIErrorMessage{Message: InvalidVisitId})
		return
	}

	err = datastore.DeleteVisit(statusId)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, nil)
}
