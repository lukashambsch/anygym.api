package datastore_test

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Status db interactions", func() {

	Describe("GetStatusList", func() {
		var statuses []models.Status

		Context("Successful call", func() {
			BeforeEach(func() {
				statuses, _ = datastore.GetStatusList()
			})

			It("should return a list of statuses", func() {
				Expect(len(statuses)).To(Equal(4))
			})
		})
	})

	Describe("GetStatus", func() {
		var status models.Status

		Context("Successful call", func() {
			var statusId int64 = 1

			BeforeEach(func() {
				status, _ = datastore.GetStatus(statusId)
			})

			It("should return the correct status", func() {
				Expect(status.StatusId).To(Equal(statusId))
			})
		})
	})

	Describe("GetStatusCount", func() {
		var count int

		Context("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetStatusCount()
			})

			It("should return the correct count", func() {
				Expect(count).To(Equal(4))
			})
		})
	})

	Describe("CreateStatus", func() {
		Context("Successful call", func() {
			It("should return the created status", func() {
				status := models.Status{StatusName: "New Status"}
				createdStatus, _ := datastore.CreateStatus(status)
				Expect(createdStatus.StatusName).To(Equal("New Status"))
			})

			It("should add a status to the db", func() {
				status := models.Status{StatusName: "New Status Two"}
				createdStatus, _ := datastore.CreateStatus(status)
				newStatus, _ := datastore.GetStatus(createdStatus.StatusId)
				Expect(newStatus.StatusName).To(Equal(createdStatus.StatusName))
			})
		})
	})

	Describe("UpdateStatus", func() {
		Context("Successful call", func() {
			It("should return the updated status", func() {
				status := models.Status{StatusName: "Updated"}
				updated, _ := datastore.UpdateStatus(status.StatusId, status)
				Expect(updated.StatusName).To(Equal("Updated"))
			})
		})
	})

	Describe("DeleteStatus", func() {
		Context("Successful call", func() {
			It("should return nil", func() {
				err := datastore.DeleteStatus(1)
				Expect(err).To(BeNil())
			})
		})
	})
})
