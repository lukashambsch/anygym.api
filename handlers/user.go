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

const UserID = "user_id"
const InvalidUserID = "Invalid " + UserID

var userFields map[string]string = map[string]string{
	"user_id":    "int",
	"email":      "string",
	"created_on": "datetime",
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userID, message := GetID(w, r, UserID)
	if message != nil {
		return
	}

	user, err := datastore.GetUser(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			WriteJSON(w, http.StatusNotFound, APIErrorMessage{Message: "Not Found"})
		} else {
			WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: err.Error()})
		}
		return
	}

	WriteJSON(w, http.StatusOK, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var statement string
	query := r.URL.Query()
	where, err := BuildWhere(userFields, query)
	if err != nil {
		WriteJSON(w, http.StatusNotFound, APIErrorMessage{Message: err.Error()})
		return
	}

	sort, err := BuildSort(userFields, query)
	if err != nil {
		WriteJSON(w, http.StatusNotFound, APIErrorMessage{Message: err.Error()})
		return
	}

	statement = fmt.Sprintf("%s %s", where, sort)
	statuses, err := datastore.GetUserList(statement)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: "Error getting user list."})
		return
	}

	WriteJSON(w, http.StatusOK, statuses)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	user := &models.User{}
	err := json.Unmarshal(body, user)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIErrorMessage{Message: err.Error()})
		return
	}

	created, err := datastore.CreateUser(*user)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: err.Error()})
		return
	}

	WriteJSON(w, http.StatusCreated, created)
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	userID, err := strconv.ParseInt(mux.Vars(r)[UserID], 10, 64)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIErrorMessage{Message: InvalidUserID})
		return
	}

	user := &models.User{}
	err = json.Unmarshal(body, user)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIErrorMessage{Message: err.Error()})
		return
	}

	updated, err := datastore.UpdateUser(userID, *user)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, updated)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(mux.Vars(r)[UserID], 10, 64)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIErrorMessage{Message: InvalidUserID})
		return
	}

	err = datastore.DeleteUser(userID)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, nil)
}
