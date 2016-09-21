package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
)

const invalidField = "Invalid field in query params."

// fields represents {field: type} mappings for db fields
func BuildWhere(fields map[string]string, params url.Values) (string, error) {
	var (
		where string = "WHERE"
		count int    = len(params)
		i     int    = 1
	)
	if count == 0 {
		return "", nil
	}

	for k, v := range params {
		if _, ok := fields[k]; ok {
			switch fields[k] {
			case "string":
				where = fmt.Sprintf("%s %s LIKE '%%%s%%'", where, k, v[0])
			case "int":
				where = fmt.Sprintf("%s %s = '%s'", where, k, v[0])
			}

			if i < count {
				where += " AND"
			}
		} else if k != "order_by" && k != "sort_order" {
			return "", fmt.Errorf(invalidField)
		}
		i += 1
	}

	if where == "WHERE" {
		return "", nil
	}

	return where, nil
}

func BuildSort(fields map[string]string, params url.Values) (string, error) {
	var (
		statement    string = "ORDER BY"
		validOrderBy bool   = false
	)

	sortOrder := params.Get("sort_order")
	orderBy := params.Get("order_by")

	if sortOrder == "" && orderBy == "" {
		return "", nil
	}

	if sortOrder == "" {
		sortOrder = "asc"
	}
	if sortOrder != "asc" && sortOrder != "desc" {
		return "", fmt.Errorf("sort_order must be either 'asc', 'desc', or ''")
	}

	for k, _ := range fields {
		if k == orderBy {
			validOrderBy = true
			break
		}
	}
	if !validOrderBy {
		return "", fmt.Errorf("Invalid order_by field.")
	}

	statement = fmt.Sprintf("%s %s %s", statement, orderBy, sortOrder)

	return statement, nil
}

func GetID(w http.ResponseWriter, r *http.Request, idField string) (int64, *APIErrorMessage) {
	id, err := strconv.ParseInt(mux.Vars(r)[idField], 10, 64)
	if err != nil {
		errMes := &APIErrorMessage{Message: "Invalid " + idField}
		WriteJSON(w, http.StatusBadRequest, errMes)
		return id, errMes
	}
	return id, nil
}

func GetNotFoundError(w http.ResponseWriter, encoder *json.Encoder, err error) {
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		encoder.Encode(APIErrorMessage{Message: "Not Found"})
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(APIErrorMessage{Message: err.Error()})
	}
}

func WriteJSON(w http.ResponseWriter, statusCode int, response interface{}) {
	// check that json can be encoded
	_, err := json.Marshal(response)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response = APIErrorMessage{Message: err.Error()}
	}

	encoder := json.NewEncoder(w)
	w.WriteHeader(statusCode)
	encoder.Encode(response)
}

func CORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
		)
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
