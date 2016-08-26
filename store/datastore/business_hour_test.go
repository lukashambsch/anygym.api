package datastore_test

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BusinessHour db interactions", func() {
	var (
		businessHourOne, businessHourTwo *models.BusinessHour
		addr                             *models.Address
		gymLocation                      *models.GymLocation
		mondayId, tuesdayId, gymId       int64 = 1, 2, 1
	)

	BeforeEach(func() {
		addr, _ = datastore.CreateAddress(models.Address{StreetAddress: "Testing"})
		gymLocation, _ = datastore.CreateGymLocation(models.GymLocation{
			GymId:        gymId,
			AddressId:    addr.AddressId,
			LocationName: "Testing",
		})
		businessHourOne, _ = datastore.CreateBusinessHour(models.BusinessHour{
			GymLocationId: gymLocation.GymLocationId,
			DayId:         &mondayId,
		})
		businessHourTwo, _ = datastore.CreateBusinessHour(models.BusinessHour{
			GymLocationId: gymLocation.GymLocationId,
			DayId:         &tuesdayId,
		})
	})

	AfterEach(func() {
		datastore.DeleteBusinessHour(businessHourOne.BusinessHourId)
		datastore.DeleteBusinessHour(businessHourTwo.BusinessHourId)
		datastore.DeleteGymLocation(gymLocation.GymLocationId)
		datastore.DeleteAddress(addr.AddressId)
	})

	Describe("GetBusinessHourList", func() {
		var businessHours []models.BusinessHour

		Describe("Successful call", func() {
			BeforeEach(func() {
				businessHours, _ = datastore.GetBusinessHourList("")
			})

			It("should return a list of businessHours", func() {
				Expect(len(businessHours)).To(Equal(2))
			})
		})
	})

	Describe("GetBusinessHour", func() {
		var businessHour *models.BusinessHour

		Describe("Successful call", func() {
			It("should return the correct businessHour", func() {
				businessHour, _ = datastore.GetBusinessHour(businessHourOne.BusinessHourId)
				Expect(businessHour.BusinessHourId).To(Equal(businessHourOne.BusinessHourId))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentId int64 = 5000
				err           error
			)

			BeforeEach(func() {
				businessHour, err = datastore.GetBusinessHour(nonExistentId)
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
				Expect(*count).To(Equal(2))
			})
		})
	})

	Describe("CreateBusinessHour", func() {
		var (
			wednesdayId  int64 = 3
			businessHour models.BusinessHour
			created      *models.BusinessHour
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				businessHour = models.BusinessHour{
					GymLocationId: gymLocation.GymLocationId,
					DayId:         &wednesdayId,
				}
				created, _ = datastore.CreateBusinessHour(businessHour)
			})

			AfterEach(func() {
				datastore.DeleteBusinessHour(created.BusinessHourId)
			})

			It("should return the created businessHour", func() {
				Expect(created.DayId).To(Equal(&wednesdayId))
			})

			It("should add a businessHour to the db", func() {
				newMember, _ := datastore.GetBusinessHour(created.BusinessHourId)
				Expect(newMember.DayId).To(Equal(&wednesdayId))
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
			wednesdayId, thursdayId int64 = 3, 4
			businessHour            models.BusinessHour
			created                 *models.BusinessHour
			updated                 *models.BusinessHour
			err                     error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				created, _ = datastore.CreateBusinessHour(models.BusinessHour{
					GymLocationId: gymLocation.GymLocationId,
					DayId:         &wednesdayId,
				})
				created.DayId = &thursdayId
				updated, _ = datastore.UpdateBusinessHour(created.BusinessHourId, *created)
			})

			AfterEach(func() {
				datastore.DeleteBusinessHour(updated.BusinessHourId)
			})

			It("should return the updated businessHour", func() {
				Expect(updated.DayId).To(Equal(&thursdayId))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				businessHour = models.BusinessHour{
					GymLocationId: gymLocation.GymLocationId,
					DayId:         &wednesdayId,
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
				var wednesdayId int64 = 3
				created, _ := datastore.CreateBusinessHour(models.BusinessHour{
					GymLocationId: gymLocation.GymLocationId,
					DayId:         &wednesdayId,
				})
				err := datastore.DeleteBusinessHour(created.BusinessHourId)
				Expect(err).To(BeNil())
			})
		})
	})
})
