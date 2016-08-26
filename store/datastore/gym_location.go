package datastore

import (
	"fmt"

	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store"
)

func GetGymLocationList(where string) ([]models.GymLocation, error) {
	var (
		gymLocations []models.GymLocation
		gymLocation  models.GymLocation
	)

	query := fmt.Sprintf("%s %s", getGymLocationListQuery, where)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&gymLocation.GymLocationId,
			&gymLocation.GymId,
			&gymLocation.AddressId,
			&gymLocation.LocationName,
			&gymLocation.PhoneNumber,
			&gymLocation.WebsiteUrl,
			&gymLocation.InNetwork,
			&gymLocation.MonthlyMemberFee,
		)
		gymLocations = append(gymLocations, gymLocation)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return gymLocations, nil
}

func GetGymLocationCount(where string) (*int, error) {
	var count int

	query := fmt.Sprintf("%s %s", getGymLocationCountQuery, where)
	row := store.DB.QueryRow(query)
	err := row.Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetGymLocation(addressId int64) (*models.GymLocation, error) {
	var gymLocation models.GymLocation

	row := store.DB.QueryRow(getGymLocationQuery, addressId)
	err := row.Scan(
		&gymLocation.GymLocationId,
		&gymLocation.GymId,
		&gymLocation.AddressId,
		&gymLocation.LocationName,
		&gymLocation.PhoneNumber,
		&gymLocation.WebsiteUrl,
		&gymLocation.InNetwork,
		&gymLocation.MonthlyMemberFee,
	)

	if err != nil {
		return nil, err
	}

	return &gymLocation, nil
}

func CreateGymLocation(gymLocation models.GymLocation) (*models.GymLocation, error) {
	var created models.GymLocation

	row := store.DB.QueryRow(
		createGymLocationQuery,
		gymLocation.GymId,
		gymLocation.AddressId,
		gymLocation.LocationName,
		gymLocation.PhoneNumber,
		gymLocation.WebsiteUrl,
		gymLocation.InNetwork,
		gymLocation.MonthlyMemberFee,
	)
	err := row.Scan(
		&created.GymLocationId,
		&created.GymId,
		&created.AddressId,
		&created.LocationName,
		&created.PhoneNumber,
		&created.WebsiteUrl,
		&created.InNetwork,
		&created.MonthlyMemberFee,
	)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateGymLocation(addressId int64, gymLocation models.GymLocation) (*models.GymLocation, error) {
	var updated models.GymLocation

	row := store.DB.QueryRow(
		updateGymLocationQuery,
		gymLocation.GymId,
		gymLocation.AddressId,
		gymLocation.LocationName,
		gymLocation.PhoneNumber,
		gymLocation.WebsiteUrl,
		gymLocation.InNetwork,
		gymLocation.MonthlyMemberFee,
		addressId,
	)
	err := row.Scan(
		&updated.GymLocationId,
		&updated.GymId,
		&updated.AddressId,
		&updated.LocationName,
		&updated.PhoneNumber,
		&updated.WebsiteUrl,
		&updated.InNetwork,
		&updated.MonthlyMemberFee,
	)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteGymLocation(addressId int64) error {
	stmt, err := store.DB.Prepare(deleteGymLocationQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(addressId)
	if err != nil {
		return err
	}

	return nil
}

const getGymLocationListQuery = `
SELECT *
FROM gym_locations
`

const getGymLocationQuery = `
SELECT *
FROM gym_locations
WHERE gym_location_id = $1
`

const createGymLocationQuery = `
INSERT INTO gym_locations (gym_id, address_id, location_name, phone_number, website_url, in_network, monthly_member_fee)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING gym_location_id, gym_id, address_id, location_name, phone_number, website_url, in_network, monthly_member_fee
`

const updateGymLocationQuery = `
UPDATE gym_locations
SET gym_id = $1, address_id = $2, location_name = $3, phone_number = $4, website_url = $5, in_network = $6, monthly_member_fee = $7
WHERE gym_location_id = $8
RETURNING gym_location_id, gym_id, address_id, location_name, phone_number, website_url, in_network, monthly_member_fee
`

const deleteGymLocationQuery = `
DELETE
FROM gym_locations
WHERE gym_location_id = $1
`

const getGymLocationCountQuery = `
SELECT count(*)
FROM gym_locations
`
