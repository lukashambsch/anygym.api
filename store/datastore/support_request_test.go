package datastore_test

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SupportRequest db interactions", func() {
	var reqOne, reqTwo *models.SupportRequest

	BeforeEach(func() {
		reqOne, _ = datastore.CreateSupportRequest(models.SupportRequest{Content: "Test Content One"})
		reqTwo, _ = datastore.CreateSupportRequest(models.SupportRequest{Content: "Test Content Two"})
	})

	AfterEach(func() {
		datastore.DeleteSupportRequest(reqOne.SupportRequestID)
		datastore.DeleteSupportRequest(reqTwo.SupportRequestID)
	})

	Describe("GetSupportRequestList", func() {
		var supportRequests []models.SupportRequest

		Describe("Successful call", func() {
			BeforeEach(func() {
				supportRequests, _ = datastore.GetSupportRequestList("")
			})

			It("should return a list of supportRequests", func() {
				Expect(len(supportRequests)).To(Equal(2))
			})
		})
	})

	Describe("GetSupportRequest", func() {
		var supportRequest *models.SupportRequest

		Describe("Successful call", func() {
			It("should return the correct supportRequest", func() {
				supportRequest, _ = datastore.GetSupportRequest(reqOne.SupportRequestID)
				Expect(supportRequest.SupportRequestID).To(Equal(reqOne.SupportRequestID))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentID int64 = 5000
				err           error
			)

			BeforeEach(func() {
				supportRequest, err = datastore.GetSupportRequest(nonExistentID)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil supportRequest", func() {
				Expect(supportRequest).To(BeNil())
			})
		})
	})

	Describe("GetSupportRequestCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetSupportRequestCount("")
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(2))
			})
		})
	})

	Describe("CreateSupportRequest", func() {
		var (
			content        string = "New SupportRequest"
			supportRequest models.SupportRequest
			created        *models.SupportRequest
            err            error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				supportRequest = models.SupportRequest{Content: content}
				created, _ = datastore.CreateSupportRequest(supportRequest)
			})

			AfterEach(func() {
				datastore.DeleteSupportRequest(created.SupportRequestID)
			})

			It("should return the created supportRequest", func() {
				Expect(created.Content).To(Equal(content))
			})

			It("should add a supportRequest to the db", func() {
				newSupportRequest, _ := datastore.GetSupportRequest(created.SupportRequestID)
				Expect(newSupportRequest.Content).To(Equal(content))
			})
		})

		Describe("Unsuccessful call", func() {
            BeforeEach(func() {
                supportRequest = models.SupportRequest{}
                _, err = datastore.CreateSupportRequest(supportRequest)
            })

            AfterEach(func() {
                datastore.DeleteSupportRequest(created.SupportRequestID)
            })

            It("should return an error object", func() {
                Expect(err).ToNot(BeNil())
            })
		})
	})

	Describe("UpdateSupportRequest", func() {
		var (
			content        string = "Fitness"
			supportRequest models.SupportRequest
			created        *models.SupportRequest
			updated        *models.SupportRequest
			err            error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				supportRequest = models.SupportRequest{Content: content}
				created, _ = datastore.CreateSupportRequest(models.SupportRequest{Content: "SupportRequest"})
				updated, _ = datastore.UpdateSupportRequest(created.SupportRequestID, supportRequest)
			})

			AfterEach(func() {
				datastore.DeleteSupportRequest(updated.SupportRequestID)
			})

			It("should return the updated supportRequest", func() {
				Expect(updated.Content).To(Equal(content))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				supportRequest = models.SupportRequest{}
				updated, err = datastore.UpdateSupportRequest(2000, supportRequest)
			})

			It("should return an error object", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil supportRequest", func() {
				Expect(updated).To(BeNil())
			})
		})
	})

	Describe("DeleteSupportRequest", func() {
		Describe("Successful call", func() {
			It("should return nil", func() {
				created, _ := datastore.CreateSupportRequest(models.SupportRequest{Content: "Test"})
				err := datastore.DeleteSupportRequest(created.SupportRequestID)
				Expect(err).To(BeNil())
			})
		})
	})
})
