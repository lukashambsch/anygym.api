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

		Describe("Successful calls", func() {
			It("should return a list of all statuses", func() {
				statuses, _ = datastore.GetStatusList("")
				Expect(len(statuses)).To(Equal(4))
			})

			It("should return a list of one partially matching status", func() {
				statuses, _ = datastore.GetStatusList("WHERE status_name LIKE '%Pend%'")
				Expect(len(statuses)).To(Equal(1))
			})

			It("should return a list of one exact match status", func() {
				statuses, _ = datastore.GetStatusList("WHERE status_name LIKE '%Pending%'")
				Expect(len(statuses)).To(Equal(1))
			})

			It("should return a list of multiple matching statuses", func() {
				statuses, _ = datastore.GetStatusList("WHERE status_name LIKE '%Denied%'")
				Expect(len(statuses)).To(Equal(2))
			})

			It("should match status_id", func() {
				statuses, _ = datastore.GetStatusList("WHERE status_id = '1'")
				Expect(len(statuses)).To(Equal(1))
			})
		})
	})

	Describe("GetStatus", func() {
		var status *models.Status

		Describe("Successful call", func() {
			var statusId int64 = 1

			It("should return the correct status", func() {
				status, _ = datastore.GetStatus(statusId)
				Expect(status.StatusId).To(Equal(statusId))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentId int64 = 5
				err           error
			)

			BeforeEach(func() {
				status, err = datastore.GetStatus(nonExistentId)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil status", func() {
				Expect(status).To(BeNil())
			})
		})
	})

	Describe("GetStatusCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetStatusCount("")
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
			err        error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				status = models.Status{StatusName: statusName}
				created, _ = datastore.CreateStatus(status)
			})

			AfterEach(func() {
				datastore.DeleteStatus(created.StatusId)
			})

			It("should return the created status", func() {
				Expect(created.StatusName).To(Equal(statusName))
			})

			It("should add a status to the db", func() {
				newStatus, _ := datastore.GetStatus(created.StatusId)
				Expect(newStatus.StatusName).To(Equal(statusName))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				status = models.Status{StatusName: "Pending"}
				created, err = datastore.CreateStatus(status)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil status", func() {
				Expect(created).To(BeNil())
			})
		})
	})

	Describe("UpdateStatus", func() {
		var (
			statusName string = "Updated"
			status     models.Status
			created    *models.Status
			updated    *models.Status
			err        error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				status = models.Status{StatusName: statusName}
				created, _ = datastore.CreateStatus(models.Status{StatusName: "Created"})
				updated, _ = datastore.UpdateStatus(created.StatusId, status)
			})

			AfterEach(func() {
				datastore.DeleteStatus(updated.StatusId)
			})

			It("should return the updated status", func() {
				Expect(updated.StatusName).To(Equal(statusName))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				status = models.Status{StatusName: "Pending"}
				updated, err = datastore.UpdateStatus(2, status)
			})

			It("should return an error object", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil status", func() {
				Expect(updated).To(BeNil())
			})
		})
	})

	Describe("DeleteStatus", func() {
		var (
			statusId int64 = 3
			status   *models.Status
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				status, _ = datastore.GetStatus(statusId)
			})

			AfterEach(func() {
				store.DB.QueryRow(
					"INSERT INTO statuses (status_id, status_name) VALUES ($1, $2)",
					status.StatusId,
					status.StatusName,
				)
			})

			It("should return nil", func() {
				err := datastore.DeleteStatus(statusId)
				Expect(err).To(BeNil())
			})
		})
	})
})
