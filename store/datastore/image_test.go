package datastore_test

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Image db interactions", func() {
	var (
		one, two, three, four *models.Image
		gymID                 int64 = 1
	)

	BeforeEach(func() {
		one, _ = datastore.CreateImage(models.Image{GymID: &gymID, ImagePath: "/tst/img/path"})
		two, _ = datastore.CreateImage(models.Image{GymID: &gymID, ImagePath: "/tst/img/path/two"})
		three, _ = datastore.CreateImage(models.Image{GymID: &gymID, ImagePath: "/tst/img/path/three"})
		four, _ = datastore.CreateImage(models.Image{GymID: &gymID, ImagePath: "/tst/img/path/four"})
	})

	AfterEach(func() {
		datastore.DeleteImage(one.ImageID)
		datastore.DeleteImage(two.ImageID)
		datastore.DeleteImage(three.ImageID)
		datastore.DeleteImage(four.ImageID)
	})

	Describe("GetImageList", func() {
		var images []models.Image

		Describe("Successful call", func() {
			BeforeEach(func() {
				images, _ = datastore.GetImageList("")
			})

			It("should return a list of images", func() {
				Expect(len(images)).To(Equal(4))
			})
		})
	})

	Describe("GetImage", func() {
		var image *models.Image

		Describe("Successful call", func() {
			It("should return the correct image", func() {
				image, _ = datastore.GetImage(one.ImageID)
				Expect(image.ImageID).To(Equal(one.ImageID))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentID int64 = 5
				err           error
			)

			BeforeEach(func() {
				image, err = datastore.GetImage(nonExistentID)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil image", func() {
				Expect(image).To(BeNil())
			})
		})
	})

	Describe("GetImageCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetImageCount("")
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(4))
			})
		})
	})

	Describe("CreateImage", func() {
		var (
			imagePath string = "/new/path"
			image     models.Image
			created   *models.Image
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				image = models.Image{GymID: &gymID, ImagePath: imagePath}
				created, _ = datastore.CreateImage(image)
			})

			AfterEach(func() {
				datastore.DeleteImage(created.ImageID)
			})

			It("should return the created image", func() {
				Expect(created.ImagePath).To(Equal(imagePath))
			})

			It("should add a image to the db", func() {
				newImage, _ := datastore.GetImage(created.ImageID)
				Expect(newImage.ImagePath).To(Equal(imagePath))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				created *models.Image
				err     error
			)

			AfterEach(func() {
				datastore.DeleteImage(created.ImageID)
			})

			It("should return an error object if user_id is not unique", func() {
				var userID int64 = int64(1)
				img := models.Image{UserID: &userID, ImagePath: "/tst/path"}
				created, _ = datastore.CreateImage(img)
				_, err = datastore.CreateImage(img)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("UpdateImage", func() {
		var (
			imagePath string = "http://urlpath.com"
			image     models.Image
			created   *models.Image
			updated   *models.Image
			err       error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				image = models.Image{GymID: &gymID, ImagePath: imagePath}
				created, _ = datastore.CreateImage(models.Image{GymID: &gymID, ImagePath: "/old/path"})
				updated, _ = datastore.UpdateImage(created.ImageID, image)
			})

			AfterEach(func() {
				datastore.DeleteImage(updated.ImageID)
			})

			It("should return the updated image", func() {
				Expect(updated.ImagePath).To(Equal(imagePath))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				image = models.Image{ImagePath: "Daily"}
				updated, err = datastore.UpdateImage(10000, image)
			})

			It("should return an error object if image to update doesn't exist", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil image", func() {
				Expect(updated).To(BeNil())
			})
		})
	})

	Describe("DeleteImage", func() {
		Describe("Successful call", func() {
			It("should return nil", func() {
				err := datastore.DeleteImage(one.ImageID)
				Expect(err).To(BeNil())
			})
		})
	})
})
