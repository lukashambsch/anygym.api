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

const StatusId = "status_id"
const InvalidStatusId = "Invalid " + StatusId

var statusFields map[string]string = map[string]string{
	"status_id":   "int",
	"status_name": "string",
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	statusId, err := strconv.ParseInt(mux.Vars(r)[StatusId], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(APIErrorMessage{Message: InvalidStatusId})
		return
	}

	status, err := datastore.GetStatus(statusId)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			encoder.Encode(APIErrorMessage{Message: "Not Found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(APIErrorMessage{Message: err.Error()})
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(status)
}

func GetStatuses(w http.ResponseWriter, r *http.Request) {
	var statement string
	encoder := json.NewEncoder(w)
	query := r.URL.Query()
	where, err := BuildWhere(statusFields, query)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		encoder.Encode(APIErrorMessage{Message: err.Error()})
		return
	}

	sort, err := BuildSort(statusFields, query)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		encoder.Encode(APIErrorMessage{Message: err.Error()})
		return
	}

	statement = fmt.Sprintf("%s %s", where, sort)

	fmt.Println(statement)
	statuses, err := datastore.GetStatusList(statement)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(APIErrorMessage{Message: "Error getting status list."})
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(statuses)
}

func PostStatus(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	encoder := json.NewEncoder(w)

	status := &models.Status{}
	err := json.Unmarshal(body, status)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(APIErrorMessage{Message: err.Error()})
		return
	}

	created, err := datastore.CreateStatus(*status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(APIErrorMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	encoder.Encode(created)
}

func PutStatus(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	encoder := json.NewEncoder(w)

	statusId, err := strconv.ParseInt(mux.Vars(r)[StatusId], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(APIErrorMessage{Message: InvalidStatusId})
		return
	}

	status := &models.Status{}
	err = json.Unmarshal(body, status)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(APIErrorMessage{Message: err.Error()})
		return
	}

	updated, err := datastore.UpdateStatus(statusId, *status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(APIErrorMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(updated)
}

func DeleteStatus(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	statusId, err := strconv.ParseInt(mux.Vars(r)[StatusId], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(APIErrorMessage{Message: InvalidStatusId})
		return
	}

	err = datastore.DeleteStatus(statusId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(APIErrorMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(nil)
}
