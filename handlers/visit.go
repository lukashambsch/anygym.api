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
	"member_id":       "int",
	"gym_location_id": "int",
	"status_id":       "int",
	"created_on":      "date",
	"modified_on":     "date",
}

func GetVisit(w http.ResponseWriter, r *http.Request) {
	visitId, message := GetId(w, r, VisitId)
	if message != nil {
		return
	}

	visit, err := datastore.GetVisit(visitId)
	if err != nil {
		if err == sql.ErrNoRows {
			WriteJSON(w, http.StatusNotFound, APIErrorMessage{Message: "Not Found"})
		} else {
			WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: err.Error()})
		}
		return
	}

	WriteJSON(w, http.StatusOK, visit)
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
	visits, err := datastore.GetVisitList(statement)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: "Error getting visit list."})
		return
	}

	WriteJSON(w, http.StatusOK, visits)
}

func PostVisit(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	visit := &models.Visit{}
	err := json.Unmarshal(body, visit)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIErrorMessage{Message: err.Error()})
		return
	}

	created, err := datastore.CreateVisit(*visit)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: err.Error()})
		return
	}

	WriteJSON(w, http.StatusCreated, created)
}

func PutVisit(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	visitId, err := strconv.ParseInt(mux.Vars(r)[VisitId], 10, 64)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIErrorMessage{Message: InvalidVisitId})
		return
	}

	visit := &models.Visit{}
	err = json.Unmarshal(body, visit)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIErrorMessage{Message: err.Error()})
		return
	}

	updated, err := datastore.UpdateVisit(visitId, *visit)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, updated)
}

func DeleteVisit(w http.ResponseWriter, r *http.Request) {
	visitId, err := strconv.ParseInt(mux.Vars(r)[VisitId], 10, 64)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIErrorMessage{Message: InvalidVisitId})
		return
	}

	err = datastore.DeleteVisit(visitId)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, nil)
}
