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

const MemberId = "member_id"
const InvalidMemberId = "Invalid " + MemberId

var memberFields map[string]string = map[string]string{
	"member_id":  "int",
	"user_id":    "int",
	"address_id": "int",
	"first_name": "string",
	"last_name":  "string",
}

func GetMember(w http.ResponseWriter, r *http.Request) {
	memberId, message := GetId(w, r, MemberId)
	if message != nil {
		return
	}

	member, err := datastore.GetMember(memberId)
	if err != nil {
		if err == sql.ErrNoRows {
			WriteJSON(w, http.StatusNotFound, APIErrorMessage{Message: "Not Found"})
		} else {
			WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: err.Error()})
		}
		return
	}

	WriteJSON(w, http.StatusOK, member)
}

func GetMembers(w http.ResponseWriter, r *http.Request) {
	var statement string
	query := r.URL.Query()
	where, err := BuildWhere(memberFields, query)
	if err != nil {
		WriteJSON(w, http.StatusNotFound, APIErrorMessage{Message: err.Error()})
		return
	}

	sort, err := BuildSort(memberFields, query)
	if err != nil {
		WriteJSON(w, http.StatusNotFound, APIErrorMessage{Message: err.Error()})
		return
	}

	statement = fmt.Sprintf("%s %s", where, sort)
	statuses, err := datastore.GetMemberList(statement)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: "Error getting member list."})
		return
	}

	WriteJSON(w, http.StatusOK, statuses)
}

func PostMember(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	member := &models.Member{}
	err := json.Unmarshal(body, member)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIErrorMessage{Message: err.Error()})
		return
	}

	created, err := datastore.CreateMember(*member)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: err.Error()})
		return
	}

	WriteJSON(w, http.StatusCreated, created)
}

func PutMember(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	memberId, err := strconv.ParseInt(mux.Vars(r)[MemberId], 10, 64)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIErrorMessage{Message: InvalidMemberId})
		return
	}

	member := &models.Member{}
	err = json.Unmarshal(body, member)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIErrorMessage{Message: err.Error()})
		return
	}

	updated, err := datastore.UpdateMember(memberId, *member)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, updated)
}

func DeleteMember(w http.ResponseWriter, r *http.Request) {
	memberId, err := strconv.ParseInt(mux.Vars(r)[MemberId], 10, 64)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIErrorMessage{Message: InvalidMemberId})
		return
	}

	err = datastore.DeleteMember(memberId)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, nil)
}
