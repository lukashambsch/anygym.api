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
		} else {
			return "", fmt.Errorf(invalidField)
		}
		i += 1
	}

	if where == "WHERE" {
		return "LIMIT 0", nil
	}

	return where, nil
}
