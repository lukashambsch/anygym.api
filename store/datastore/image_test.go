package datastore_test

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Image db interactions", func() {
	var one, two, three, four *models.Image

	BeforeEach(func() {
		one, _ = datastore.CreateImage(models.Image{ImagePath: "/tst/img/path"})
		two, _ = datastore.CreateImage(models.Image{ImagePath: "/tst/img/path/two"})
		three, _ = datastore.CreateImage(models.Image{ImagePath: "/tst/img/path/three"})
		four, _ = datastore.CreateImage(models.Image{ImagePath: "/tst/img/path/four"})
	})

	AfterEach(func() {
		datastore.DeleteImage(one.ImageId)
		datastore.DeleteImage(two.ImageId)
		datastore.DeleteImage(three.ImageId)
		datastore.DeleteImage(four.ImageId)
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
				image, _ = datastore.GetImage(one.ImageId)
				Expect(image.ImageId).To(Equal(one.ImageId))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentId int64 = 5
				err           error
			)

			BeforeEach(func() {
				image, err = datastore.GetImage(nonExistentId)
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
				image = models.Image{ImagePath: imagePath}
				created, _ = datastore.CreateImage(image)
			})

			AfterEach(func() {
				datastore.DeleteImage(created.ImageId)
			})

			It("should return the created image", func() {
				Expect(created.ImagePath).To(Equal(imagePath))
			})

			It("should add a image to the db", func() {
				newImage, _ := datastore.GetImage(created.ImageId)
				Expect(newImage.ImagePath).To(Equal(imagePath))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				created *models.Image
				err     error
			)

			AfterEach(func() {
				datastore.DeleteImage(created.ImageId)
			})

			It("should return an error object if user_id is not unique", func() {
				var userId int64 = int64(1)
				img := models.Image{UserId: &userId, ImagePath: "/tst/path"}
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
				image = models.Image{ImagePath: imagePath}
				created, _ = datastore.CreateImage(models.Image{ImagePath: "/old/path"})
				updated, _ = datastore.UpdateImage(created.ImageId, image)
			})

			AfterEach(func() {
				datastore.DeleteImage(updated.ImageId)
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
				err := datastore.DeleteImage(one.ImageId)
				Expect(err).To(BeNil())
			})
		})
	})
})
