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
		err = rows.Scan(&image.ImageID, &image.GymID, &image.GymLocationID, &image.UserID, &image.ImagePath)
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

func GetImage(imageID int64) (*models.Image, error) {
	var image models.Image

	row := store.DB.QueryRow(getImageQuery, imageID)
	err := row.Scan(&image.ImageID, &image.GymID, &image.GymLocationID, &image.UserID, &image.ImagePath)
	if err != nil {
		return nil, err
	}

	return &image, nil
}

func CreateImage(image models.Image) (*models.Image, error) {
	var created models.Image

	row := store.DB.QueryRow(createImageQuery, image.GymID, image.GymLocationID, image.UserID, image.ImagePath)
	err := row.Scan(&created.ImageID, &created.GymID, &created.GymLocationID, &created.UserID, &created.ImagePath)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateImage(imageID int64, image models.Image) (*models.Image, error) {
	var updated models.Image

	row := store.DB.QueryRow(updateImageQuery, image.GymID, image.GymLocationID, image.UserID, image.ImagePath, imageID)
	err := row.Scan(&updated.ImageID, &updated.GymID, &updated.GymLocationID, &updated.UserID, &updated.ImagePath)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteImage(imageID int64) error {
	stmt, err := store.DB.Prepare(deleteImageQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(imageID)
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
