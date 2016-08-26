package datastore

import (
	"fmt"

	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store"
)

func GetDeviceList(where string) ([]models.Device, error) {
	var (
		devices []models.Device
		device  models.Device
	)

	query := fmt.Sprintf("%s %s", getDeviceListQuery, where)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&device.DeviceId, &device.UserId, &device.DeviceToken)
		devices = append(devices, device)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return devices, nil
}

func GetDeviceCount(where string) (*int, error) {
	var count int

	query := fmt.Sprintf("%s %s", getDeviceCountQuery, where)
	row := store.DB.QueryRow(query)
	err := row.Scan(&count)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetDevice(deviceId int64) (*models.Device, error) {
	var device models.Device

	row := store.DB.QueryRow(getDeviceQuery, deviceId)
	err := row.Scan(&device.DeviceId, &device.UserId, &device.DeviceToken)
	if err != nil {
		return nil, err
	}

	return &device, nil
}

func CreateDevice(device models.Device) (*models.Device, error) {
	var created models.Device

	row := store.DB.QueryRow(createDeviceQuery, device.UserId, device.DeviceToken)
	err := row.Scan(&created.DeviceId, &created.UserId, &created.DeviceToken)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateDevice(deviceId int64, device models.Device) (*models.Device, error) {
	var updated models.Device

	row := store.DB.QueryRow(updateDeviceQuery, device.UserId, device.DeviceToken, deviceId)
	err := row.Scan(&updated.DeviceId, &updated.UserId, &updated.DeviceToken)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteDevice(deviceId int64) error {
	stmt, err := store.DB.Prepare(deleteDeviceQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(deviceId)
	if err != nil {
		return err
	}

	return nil
}

const getDeviceListQuery = `
SELECT *
FROM devices
`

const getDeviceQuery = `
SELECT *
FROM devices
WHERE device_id = $1
`

const createDeviceQuery = `
INSERT INTO devices (user_id, device_token)
VALUES ($1, $2)
RETURNING device_id, user_id, device_token
`

const updateDeviceQuery = `
UPDATE devices
SET user_id = $1, device_token = $2
WHERE device_id = $3
RETURNING device_id, user_id, device_token
`

const deleteDeviceQuery = `
DELETE
FROM devices
WHERE device_id = $1
`

const getDeviceCountQuery = `
SELECT count(*)
FROM devices
`
