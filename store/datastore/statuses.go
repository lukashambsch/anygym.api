package datastore

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store"
)

func GetStatusList() ([]models.Status, error) {
	var (
		statuses []models.Status
		status   models.Status
	)

	db := store.Open()

	rows, err := db.Query("SELECT * FROM statuses;")
	if err != nil {
		return statuses, err
	}

	for rows.Next() {
		err = rows.Scan(&status.StatusId, &status.StatusName)
		statuses = append(statuses, status)
		if err != nil {
			return statuses, err
		}
	}
	defer rows.Close()

	return statuses, nil
}
