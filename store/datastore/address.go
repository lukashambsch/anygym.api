package datastore

import (
	"fmt"

	"github.com/lukashambsch/anygym.api/models"
	"github.com/lukashambsch/anygym.api/store"
)

func GetAddressList(where string) ([]models.Address, error) {
	var (
		addresses []models.Address
		address   models.Address
	)

	query := fmt.Sprintf("%s %s", getAddressListQuery, where)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&address.AddressID,
			&address.Country,
			&address.StateRegion,
			&address.City,
			&address.PostalArea,
			&address.StreetAddress,
			&address.Latitude,
			&address.Longitude,
		)
		addresses = append(addresses, address)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return addresses, nil
}

func GetAddressCount(where string) (*int, error) {
	var count int

	query := fmt.Sprintf("%s %s", getAddressCountQuery, where)
	row := store.DB.QueryRow(query)
	err := row.Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetAddress(addressID int64) (*models.Address, error) {
	var address models.Address

	row := store.DB.QueryRow(getAddressQuery, addressID)
	err := row.Scan(
		&address.AddressID,
		&address.Country,
		&address.StateRegion,
		&address.City,
		&address.PostalArea,
		&address.StreetAddress,
		&address.Latitude,
		&address.Longitude,
	)

	if err != nil {
		return nil, err
	}

	return &address, nil
}

func CreateAddress(address models.Address) (*models.Address, error) {
	var created models.Address

	row := store.DB.QueryRow(
		createAddressQuery,
		address.Country,
		address.StateRegion,
		address.City,
		address.PostalArea,
		address.StreetAddress,
		address.Latitude,
		address.Longitude,
	)
	err := row.Scan(
		&created.AddressID,
		&created.Country,
		&created.StateRegion,
		&created.City,
		&created.PostalArea,
		&created.StreetAddress,
		&created.Latitude,
		&created.Longitude,
	)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateAddress(addressID int64, address models.Address) (*models.Address, error) {
	var updated models.Address

	row := store.DB.QueryRow(
		updateAddressQuery,
		address.Country,
		address.StateRegion,
		address.City,
		address.PostalArea,
		address.StreetAddress,
		address.Latitude,
		address.Longitude,
		addressID,
	)
	err := row.Scan(
		&updated.AddressID,
		&updated.Country,
		&updated.StateRegion,
		&updated.City,
		&updated.PostalArea,
		&updated.StreetAddress,
		&updated.Latitude,
		&updated.Longitude,
	)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteAddress(addressID int64) error {
	stmt, err := store.DB.Prepare(deleteAddressQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(addressID)
	if err != nil {
		return err
	}

	return nil
}

const getAddressListQuery = `
SELECT *
FROM addresses
`

const getAddressQuery = `
SELECT *
FROM addresses
WHERE address_id = $1
`

const createAddressQuery = `
INSERT INTO addresses (country, state_region, city, postal_area, street_address, latitude, longitude)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING address_id, country, state_region, city, postal_area, street_address, latitude, longitude
`

const updateAddressQuery = `
UPDATE addresses
SET country = $1, state_region = $2, city = $3, postal_area = $4, street_address = $5, latitude = $6, longitude = $7
WHERE address_id = $8
RETURNING address_id, country, state_region, city, postal_area, street_address, latitude, longitude
`

const deleteAddressQuery = `
DELETE
FROM addresses
WHERE address_id = $1
`

const getAddressCountQuery = `
SELECT count(*)
FROM addresses
`
