package datastore_test

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store"
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
		var status *models.Status

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
		var count *int

		Context("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetStatusCount()
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(4))
			})
		})
	})

	Describe("CreateStatus", func() {
		var (
			statusName string = "New Status"
			status     models.Status
			created    *models.Status
		)

		BeforeEach(func() {
			status = models.Status{StatusName: statusName}
			created, _ = datastore.CreateStatus(status)
		})

		Context("Successful call", func() {
			It("should return the created status", func() {
				Expect(created.StatusName).To(Equal(statusName))
			})

			It("should add a status to the db", func() {
				newStatus, _ := datastore.GetStatus(created.StatusId)
				Expect(newStatus.StatusName).To(Equal(statusName))
			})
		})

		AfterEach(func() {
			datastore.DeleteStatus(created.StatusId)
		})
	})

	Describe("UpdateStatus", func() {
		var (
			statusName string = "Updated"
			status     models.Status
			created    *models.Status
			updated    *models.Status
		)

		BeforeEach(func() {
			status = models.Status{StatusName: statusName}
			created, _ = datastore.CreateStatus(models.Status{StatusName: "Created"})
			updated, _ = datastore.UpdateStatus(created.StatusId, status)
		})

		Context("Successful call", func() {
			It("should return the updated status", func() {
				Expect(updated.StatusName).To(Equal(statusName))
			})
		})

		AfterEach(func() {
			datastore.DeleteStatus(updated.StatusId)
		})
	})

	Describe("DeleteStatus", func() {
		var (
			statusId int64 = 1
			status   *models.Status
		)

		BeforeEach(func() {
			status, _ = datastore.GetStatus(statusId)
		})

		Context("Successful call", func() {
			It("should return nil", func() {
				err := datastore.DeleteStatus(statusId)
				Expect(err).To(BeNil())
			})
		})

		AfterEach(func() {
			store.DB.QueryRow(
				"INSERT INTO statuses (status_id, status_name) VALUES ($1, $2)",
				status.StatusId,
				status.StatusName,
			)
		})
	})
})
