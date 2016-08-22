package handlers

import (
	"fmt"
	"net/url"
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
