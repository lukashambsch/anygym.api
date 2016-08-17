package datastore

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store"
)

func GetAddressList() ([]models.Address, error) {
	var (
		addresses []models.Address
		address   models.Address
	)

	rows, err := store.DB.Query(getAddressListQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&address.AddressId,
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

func GetAddressCount() (*int, error) {
	var count int

	row := store.DB.QueryRow(getAddressCountQuery)
	err := row.Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetAddress(addressId int64) (*models.Address, error) {
	var address models.Address

	row := store.DB.QueryRow(getAddressQuery, addressId)
	err := row.Scan(
		&address.AddressId,
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
		&created.AddressId,
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

func UpdateAddress(addressId int64, address models.Address) (*models.Address, error) {
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
		addressId,
	)
	err := row.Scan(
		&updated.AddressId,
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

func DeleteAddress(addressId int64) error {
	stmt, err := store.DB.Prepare(deleteAddressQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(addressId)
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
