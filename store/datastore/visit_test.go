package datastore_test

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Visit db interactions", func() {
	var (
		visitOne, visitTwo *models.Visit
		addr               *models.Address
		gymLocation        *models.GymLocation
		gymID              int64 = 1
		memberID           int64 = 1
		statusID           int64 = 1
	)

	BeforeEach(func() {
		addr, _ = datastore.CreateAddress(models.Address{StreetAddress: "Testing"})
		gymLocation, _ = datastore.CreateGymLocation(models.GymLocation{
			GymID:        gymID,
			AddressID:    addr.AddressID,
			LocationName: "Testing",
		})
		visitOne, _ = datastore.CreateVisit(models.Visit{
			MemberID:      memberID,
			GymLocationID: gymLocation.GymLocationID,
			StatusID:      statusID,
		})
		visitTwo, _ = datastore.CreateVisit(models.Visit{
			MemberID:      memberID,
			GymLocationID: gymLocation.GymLocationID,
			StatusID:      statusID,
		})
	})

	AfterEach(func() {
		datastore.DeleteVisit(visitOne.VisitID)
		datastore.DeleteVisit(visitTwo.VisitID)
		datastore.DeleteGymLocation(gymLocation.GymLocationID)
		datastore.DeleteAddress(addr.AddressID)
	})

	Describe("GetVisitList", func() {
		var visits []models.Visit

		Describe("Successful call", func() {
			BeforeEach(func() {
				visits, _ = datastore.GetVisitList("")
			})

			It("should return a list of visits", func() {
				Expect(len(visits)).To(Equal(7))
			})
		})
	})

	Describe("GetVisit", func() {
		var visit *models.Visit

		Describe("Successful call", func() {
			It("should return the correct visit", func() {
				visit, _ = datastore.GetVisit(visitOne.VisitID)
				Expect(visit.VisitID).To(Equal(visitOne.VisitID))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentID int64 = 5000
				err           error
			)

			BeforeEach(func() {
				visit, err = datastore.GetVisit(nonExistentID)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil visit", func() {
				Expect(visit).To(BeNil())
			})
		})
	})

	Describe("GetVisitCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetVisitCount("")
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(7))
			})
		})
	})

	Describe("CreateVisit", func() {
		var (
			visit   models.Visit
			created *models.Visit
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				visit = models.Visit{
					MemberID:      memberID,
					GymLocationID: gymLocation.GymLocationID,
					StatusID:      statusID,
				}
				created, _ = datastore.CreateVisit(visit)
			})

			AfterEach(func() {
				datastore.DeleteVisit(created.VisitID)
			})

			It("should return the created visit", func() {
				Expect(created.StatusID).To(Equal(statusID))
			})

			It("should add a visit to the db", func() {
				newMember, _ := datastore.GetVisit(created.VisitID)
				Expect(newMember.StatusID).To(Equal(statusID))
			})
		})

		Describe("Unsuccessful call", func() {
			It("should return an error object if no fk ids", func() {
				v := models.Visit{}
				_, err := datastore.CreateVisit(v)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("UpdateVisit", func() {
		var (
			otherID int64 = 3
			visit   models.Visit
			created *models.Visit
			updated *models.Visit
			err     error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				created, _ = datastore.CreateVisit(models.Visit{
					MemberID:      memberID,
					GymLocationID: gymLocation.GymLocationID,
					StatusID:      statusID,
				})
				created.StatusID = otherID
				updated, _ = datastore.UpdateVisit(created.VisitID, *created)
			})

			AfterEach(func() {
				datastore.DeleteVisit(updated.VisitID)
			})

			It("should return the updated visit", func() {
				Expect(updated.StatusID).To(Equal(otherID))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				visit = models.Visit{
					GymLocationID: gymLocation.GymLocationID,
					StatusID:      statusID,
				}
				updated, err = datastore.UpdateVisit(5000, visit)
			})

			It("should return an error object if visit to update doesn't exist", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil visit", func() {
				Expect(updated).To(BeNil())
			})
		})
	})

	Describe("DeleteVisit", func() {
		Describe("Successful call", func() {
			It("should return nil", func() {
				created, _ := datastore.CreateVisit(models.Visit{
					MemberID:      memberID,
					GymLocationID: gymLocation.GymLocationID,
					StatusID:      statusID,
				})
				err := datastore.DeleteVisit(created.VisitID)
				Expect(err).To(BeNil())
			})
		})
	})
})
