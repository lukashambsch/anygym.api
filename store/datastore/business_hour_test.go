package datastore_test

import (
	"github.com/lukashambsch/anygym.api/models"
	"github.com/lukashambsch/anygym.api/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BusinessHour db interactions", func() {
	var (
		businessHourOne, businessHourTwo *models.BusinessHour
		addr                             *models.Address
		gymLocation                      *models.GymLocation
		mondayID, tuesdayID, gymID       int64 = 1, 2, 1
	)

	BeforeEach(func() {
		addr, _ = datastore.CreateAddress(models.Address{StreetAddress: "Testing"})
		gymLocation, _ = datastore.CreateGymLocation(models.GymLocation{
			GymID:        gymID,
			AddressID:    addr.AddressID,
			LocationName: "Testing",
		})
		businessHourOne, _ = datastore.CreateBusinessHour(models.BusinessHour{
			GymLocationID: gymLocation.GymLocationID,
			DayID:         &mondayID,
		})
		businessHourTwo, _ = datastore.CreateBusinessHour(models.BusinessHour{
			GymLocationID: gymLocation.GymLocationID,
			DayID:         &tuesdayID,
		})
	})

	AfterEach(func() {
		datastore.DeleteBusinessHour(businessHourOne.BusinessHourID)
		datastore.DeleteBusinessHour(businessHourTwo.BusinessHourID)
		datastore.DeleteGymLocation(gymLocation.GymLocationID)
		datastore.DeleteAddress(addr.AddressID)
	})

	Describe("GetBusinessHourList", func() {
		var businessHours []models.BusinessHour

		Describe("Successful call", func() {
			BeforeEach(func() {
				businessHours, _ = datastore.GetBusinessHourList("")
			})

			It("should return a list of businessHours", func() {
				Expect(len(businessHours)).To(Equal(16))
			})
		})
	})

	Describe("GetBusinessHour", func() {
		var businessHour *models.BusinessHour

		Describe("Successful call", func() {
			It("should return the correct businessHour", func() {
				businessHour, _ = datastore.GetBusinessHour(businessHourOne.BusinessHourID)
				Expect(businessHour.BusinessHourID).To(Equal(businessHourOne.BusinessHourID))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentID int64 = 5000
				err           error
			)

			BeforeEach(func() {
				businessHour, err = datastore.GetBusinessHour(nonExistentID)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil businessHour", func() {
				Expect(businessHour).To(BeNil())
			})
		})
	})

	Describe("GetBusinessHourCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetBusinessHourCount("")
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(16))
			})
		})
	})

	Describe("CreateBusinessHour", func() {
		var (
			wednesdayID  int64 = 3
			businessHour models.BusinessHour
			created      *models.BusinessHour
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				businessHour = models.BusinessHour{
					GymLocationID: gymLocation.GymLocationID,
					DayID:         &wednesdayID,
				}
				created, _ = datastore.CreateBusinessHour(businessHour)
			})

			AfterEach(func() {
				datastore.DeleteBusinessHour(created.BusinessHourID)
			})

			It("should return the created businessHour", func() {
				Expect(created.DayID).To(Equal(&wednesdayID))
			})

			It("should add a businessHour to the db", func() {
				newMember, _ := datastore.GetBusinessHour(created.BusinessHourID)
				Expect(newMember.DayID).To(Equal(&wednesdayID))
			})
		})

		Describe("Unsuccessful call", func() {
			It("should return an error object if no day_id or holiday_id", func() {
				mbr := models.BusinessHour{}
				_, err := datastore.CreateBusinessHour(mbr)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("UpdateBusinessHour", func() {
		var (
			wednesdayID, thursdayID int64 = 3, 4
			businessHour            models.BusinessHour
			created                 *models.BusinessHour
			updated                 *models.BusinessHour
			err                     error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				created, _ = datastore.CreateBusinessHour(models.BusinessHour{
					GymLocationID: gymLocation.GymLocationID,
					DayID:         &wednesdayID,
				})
				created.DayID = &thursdayID
				updated, _ = datastore.UpdateBusinessHour(created.BusinessHourID, *created)
			})

			AfterEach(func() {
				datastore.DeleteBusinessHour(updated.BusinessHourID)
			})

			It("should return the updated businessHour", func() {
				Expect(updated.DayID).To(Equal(&thursdayID))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				businessHour = models.BusinessHour{
					GymLocationID: gymLocation.GymLocationID,
					DayID:         &wednesdayID,
				}
				updated, err = datastore.UpdateBusinessHour(5000, businessHour)
			})

			It("should return an error object if businessHour to update doesn't exist", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil businessHour", func() {
				Expect(updated).To(BeNil())
			})
		})
	})

	Describe("DeleteBusinessHour", func() {
		Describe("Successful call", func() {
			It("should return nil", func() {
				var wednesdayID int64 = 3
				created, _ := datastore.CreateBusinessHour(models.BusinessHour{
					GymLocationID: gymLocation.GymLocationID,
					DayID:         &wednesdayID,
				})
				err := datastore.DeleteBusinessHour(created.BusinessHourID)
				Expect(err).To(BeNil())
			})
		})
	})
})
