package datastore

import (
	"fmt"

	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store"
)

func GetImageList(where string) ([]models.Image, error) {
	var (
		images []models.Image
		image  models.Image
	)

	query := fmt.Sprintf("%s %s", getImageListQuery, where)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&image.ImageId, &image.GymId, &image.GymLocationId, &image.UserId, &image.ImagePath)
		images = append(images, image)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return images, nil
}

func GetImageCount(where string) (*int, error) {
	var count int

	query := fmt.Sprintf("%s %s", getImageCountQuery, where)
	row := store.DB.QueryRow(query)
	err := row.Scan(&count)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetImage(imageId int64) (*models.Image, error) {
	var image models.Image

	row := store.DB.QueryRow(getImageQuery, imageId)
	err := row.Scan(&image.ImageId, &image.GymId, &image.GymLocationId, &image.UserId, &image.ImagePath)
	if err != nil {
		return nil, err
	}

	return &image, nil
}

func CreateImage(image models.Image) (*models.Image, error) {
	var created models.Image

	row := store.DB.QueryRow(createImageQuery, image.GymId, image.GymLocationId, image.UserId, image.ImagePath)
	err := row.Scan(&created.ImageId, &created.GymId, &created.GymLocationId, &created.UserId, &created.ImagePath)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateImage(imageId int64, image models.Image) (*models.Image, error) {
	var updated models.Image

	row := store.DB.QueryRow(updateImageQuery, image.GymId, image.GymLocationId, image.UserId, image.ImagePath, imageId)
	err := row.Scan(&updated.ImageId, &updated.GymId, &updated.GymLocationId, &updated.UserId, &updated.ImagePath)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteImage(imageId int64) error {
	stmt, err := store.DB.Prepare(deleteImageQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(imageId)
	if err != nil {
		return err
	}

	return nil
}

const getImageListQuery = `
SELECT *
FROM images
`

const getImageQuery = `
SELECT *
FROM images
WHERE image_id = $1
`

const createImageQuery = `
INSERT INTO images (gym_id, gym_location_id, user_id, image_path)
VALUES ($1, $2, $3, $4)
RETURNING image_id, gym_id, gym_location_id, user_id, image_path
`

const updateImageQuery = `
UPDATE images
SET gym_id = $1, gym_location_id = $2, user_id = $3, image_path = $4
WHERE image_id = $5
RETURNING image_id, gym_id, gym_location_id, user_id, image_path
`

const deleteImageQuery = `
DELETE
FROM images
WHERE image_id = $1
`

const getImageCountQuery = `
SELECT count(*)
FROM images
`
