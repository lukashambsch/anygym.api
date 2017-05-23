package datastore

import (
	"fmt"

	"github.com/lukashambsch/anygym.api/models"
	"github.com/lukashambsch/anygym.api/store"
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
			&gymLocation.GymLocationID,
			&gymLocation.GymID,
			&gymLocation.AddressID,
			&gymLocation.LocationName,
			&gymLocation.PhoneNumber,
			&gymLocation.WebsiteUrl,
			&gymLocation.InNetwork,
			&gymLocation.MonthlyMemberFee,
			&gymLocation.Address.AddressID,
			&gymLocation.Address.Country,
			&gymLocation.Address.StateRegion,
			&gymLocation.Address.City,
			&gymLocation.Address.PostalArea,
			&gymLocation.Address.StreetAddress,
			&gymLocation.Address.Latitude,
			&gymLocation.Address.Longitude,
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

func GetGymLocation(addressID int64) (*models.GymLocation, error) {
	var gymLocation models.GymLocation

	row := store.DB.QueryRow(getGymLocationQuery, addressID)
	err := row.Scan(
		&gymLocation.GymLocationID,
		&gymLocation.GymID,
		&gymLocation.AddressID,
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
		gymLocation.GymID,
		gymLocation.AddressID,
		gymLocation.LocationName,
		gymLocation.PhoneNumber,
		gymLocation.WebsiteUrl,
		gymLocation.InNetwork,
		gymLocation.MonthlyMemberFee,
	)
	err := row.Scan(
		&created.GymLocationID,
		&created.GymID,
		&created.AddressID,
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

func UpdateGymLocation(addressID int64, gymLocation models.GymLocation) (*models.GymLocation, error) {
	var updated models.GymLocation

	row := store.DB.QueryRow(
		updateGymLocationQuery,
		gymLocation.GymID,
		gymLocation.AddressID,
		gymLocation.LocationName,
		gymLocation.PhoneNumber,
		gymLocation.WebsiteUrl,
		gymLocation.InNetwork,
		gymLocation.MonthlyMemberFee,
		addressID,
	)
	err := row.Scan(
		&updated.GymLocationID,
		&updated.GymID,
		&updated.AddressID,
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

func DeleteGymLocation(addressID int64) error {
	stmt, err := store.DB.Prepare(deleteGymLocationQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(addressID)
	if err != nil {
		return err
	}

	return nil
}

const getGymLocationListQuery = `
SELECT
    gl.gym_location_id,
    gl.gym_id,
    gl.address_id,
    gl.location_name,
    gl.phone_number,
    gl.website_url,
    gl.in_network,
    gl.monthly_member_fee,
    a.address_id,
    a.country,
    a.state_region,
    a.city,
    a.postal_area,
    a.street_address,
    a.latitude,
    a.longitude
FROM gym_locations AS gl
JOIN addresses AS a
ON gl.address_id = a.address_id
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
