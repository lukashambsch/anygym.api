package datastore_test

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SupportSource db interactions", func() {
	var supportSourceId int64 = 1

	Describe("GetSupportSourceList", func() {
		var supportSources []models.SupportSource

		Describe("Successful call", func() {
			BeforeEach(func() {
				supportSources, _ = datastore.GetSupportSourceList("")
			})

			It("should return a list of supportSources", func() {
				Expect(len(supportSources)).To(Equal(6))
			})
		})
	})

	Describe("GetSupportSource", func() {
		var supportSource *models.SupportSource

		Describe("Successful call", func() {
			It("should return the correct supportSource", func() {
				supportSource, _ = datastore.GetSupportSource(supportSourceId)
				Expect(supportSource.SupportSourceId).To(Equal(supportSourceId))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentId int64 = 5000
				err           error
			)

			BeforeEach(func() {
				supportSource, err = datastore.GetSupportSource(nonExistentId)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil supportSource", func() {
				Expect(supportSource).To(BeNil())
			})
		})
	})

	Describe("GetSupportSourceCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetSupportSourceCount("")
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(6))
			})
		})
	})

	Describe("CreateSupportSource", func() {
		var (
			supportSourceName string = "New SupportSource"
			supportSource     models.SupportSource
			created           *models.SupportSource
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				supportSource = models.SupportSource{SupportSourceName: supportSourceName}
				created, _ = datastore.CreateSupportSource(supportSource)
			})

			AfterEach(func() {
				datastore.DeleteSupportSource(created.SupportSourceId)
			})

			It("should return the created supportSource", func() {
				Expect(created.SupportSourceName).To(Equal(supportSourceName))
			})

			It("should add a supportSource to the db", func() {
				newSupportSource, _ := datastore.GetSupportSource(created.SupportSourceId)
				Expect(newSupportSource.SupportSourceName).To(Equal(supportSourceName))
			})
		})

		Describe("Unsuccessful call", func() {
			var created *models.SupportSource

			AfterEach(func() {
				datastore.DeleteSupportSource(created.SupportSourceId)
			})

			It("should return an error object if supportSource is not unique", func() {
				name := "Test Name"
				pln := models.SupportSource{SupportSourceName: name}
				created, _ = datastore.CreateSupportSource(pln)
				_, err := datastore.CreateSupportSource(pln)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("UpdateSupportSource", func() {
		var (
			supportSourceName string = "Anytime"
			supportSource     models.SupportSource
			created           *models.SupportSource
			updated           *models.SupportSource
			err               error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				supportSource = models.SupportSource{SupportSourceName: supportSourceName}
				created, _ = datastore.CreateSupportSource(models.SupportSource{SupportSourceName: "Daily"})
				updated, _ = datastore.UpdateSupportSource(created.SupportSourceId, supportSource)
			})

			AfterEach(func() {
				datastore.DeleteSupportSource(updated.SupportSourceId)
			})

			It("should return the updated supportSource", func() {
				Expect(updated.SupportSourceName).To(Equal(supportSourceName))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				supportSource = models.SupportSource{SupportSourceName: "Daily"}
				updated, err = datastore.UpdateSupportSource(10000, supportSource)
			})

			It("should return an error object if supportSource to update doesn't exist", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil supportSource", func() {
				Expect(updated).To(BeNil())
			})
		})
	})

	Describe("DeleteSupportSource", func() {
		Describe("Successful call", func() {
			It("should return nil", func() {
				created, _ := datastore.CreateSupportSource(models.SupportSource{SupportSourceName: "Testing"})
				err := datastore.DeleteSupportSource(created.SupportSourceId)
				Expect(err).To(BeNil())
			})
		})
	})
})
