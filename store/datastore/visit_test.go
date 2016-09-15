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
		gymId              int64 = 1
		memberId           int64 = 1
		statusId           int64 = 1
	)

	BeforeEach(func() {
		addr, _ = datastore.CreateAddress(models.Address{StreetAddress: "Testing"})
		gymLocation, _ = datastore.CreateGymLocation(models.GymLocation{
			GymId:        gymId,
			AddressId:    addr.AddressId,
			LocationName: "Testing",
		})
		visitOne, _ = datastore.CreateVisit(models.Visit{
			MemberId:      memberId,
			GymLocationId: gymLocation.GymLocationId,
			StatusId:      statusId,
		})
		visitTwo, _ = datastore.CreateVisit(models.Visit{
			MemberId:      memberId,
			GymLocationId: gymLocation.GymLocationId,
			StatusId:      statusId,
		})
	})

	AfterEach(func() {
		datastore.DeleteVisit(visitOne.VisitId)
		datastore.DeleteVisit(visitTwo.VisitId)
		datastore.DeleteGymLocation(gymLocation.GymLocationId)
		datastore.DeleteAddress(addr.AddressId)
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
				visit, _ = datastore.GetVisit(visitOne.VisitId)
				Expect(visit.VisitId).To(Equal(visitOne.VisitId))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentId int64 = 5000
				err           error
			)

			BeforeEach(func() {
				visit, err = datastore.GetVisit(nonExistentId)
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
					MemberId:      memberId,
					GymLocationId: gymLocation.GymLocationId,
					StatusId:      statusId,
				}
				created, _ = datastore.CreateVisit(visit)
			})

			AfterEach(func() {
				datastore.DeleteVisit(created.VisitId)
			})

			It("should return the created visit", func() {
				Expect(created.StatusId).To(Equal(statusId))
			})

			It("should add a visit to the db", func() {
				newMember, _ := datastore.GetVisit(created.VisitId)
				Expect(newMember.StatusId).To(Equal(statusId))
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
			otherId int64 = 3
			visit   models.Visit
			created *models.Visit
			updated *models.Visit
			err     error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				created, _ = datastore.CreateVisit(models.Visit{
					MemberId:      memberId,
					GymLocationId: gymLocation.GymLocationId,
					StatusId:      statusId,
				})
				created.StatusId = otherId
				updated, _ = datastore.UpdateVisit(created.VisitId, *created)
			})

			AfterEach(func() {
				datastore.DeleteVisit(updated.VisitId)
			})

			It("should return the updated visit", func() {
				Expect(updated.StatusId).To(Equal(otherId))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				visit = models.Visit{
					GymLocationId: gymLocation.GymLocationId,
					StatusId:      statusId,
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
					MemberId:      memberId,
					GymLocationId: gymLocation.GymLocationId,
					StatusId:      statusId,
				})
				err := datastore.DeleteVisit(created.VisitId)
				Expect(err).To(BeNil())
			})
		})
	})
})
