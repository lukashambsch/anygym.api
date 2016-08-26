package datastore_test

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GymLocation db interactions", func() {
	var (
		one, two      *models.GymLocation
		addr, address *models.Address
	)

	BeforeEach(func() {
		addr, _ = datastore.CreateAddress(models.Address{StreetAddress: "Testing"})
		address, _ = datastore.CreateAddress(models.Address{StreetAddress: "Testing Two"})
		one, _ = datastore.CreateGymLocation(models.GymLocation{
			AddressId:    addr.AddressId,
			LocationName: "Testing",
		})
		two, _ = datastore.CreateGymLocation(models.GymLocation{
			AddressId:    address.AddressId,
			LocationName: "Testing Two",
		})
	})

	AfterEach(func() {
		datastore.DeleteGymLocation(one.GymLocationId)
		datastore.DeleteGymLocation(two.GymLocationId)
		datastore.DeleteAddress(addr.AddressId)
		datastore.DeleteAddress(address.AddressId)
	})

	Describe("GetGymLocationList", func() {
		var gymLocations []models.GymLocation

		Describe("Successful call", func() {
			BeforeEach(func() {
				gymLocations, _ = datastore.GetGymLocationList("")
			})

			It("should return a list of gymLocations", func() {
				Expect(len(gymLocations)).To(Equal(2))
			})
		})
	})

	Describe("GetGymLocation", func() {
		var gymLocation *models.GymLocation

		Describe("Successful call", func() {
			It("should return the correct gymLocation", func() {
				gymLocation, _ = datastore.GetGymLocation(one.GymLocationId)
				Expect(gymLocation.GymLocationId).To(Equal(one.GymLocationId))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentId int64 = 5000
				err           error
			)

			BeforeEach(func() {
				gymLocation, err = datastore.GetGymLocation(nonExistentId)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil gymLocation", func() {
				Expect(gymLocation).To(BeNil())
			})
		})
	})

	Describe("GetGymLocationCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetGymLocationCount("")
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(2))
			})
		})
	})

	Describe("CreateGymLocation", func() {
		var (
			locationName string = "New GymLocation"
			gymLocation  models.GymLocation
			created      *models.GymLocation
			newAddr      *models.Address
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				newAddr, _ = datastore.CreateAddress(models.Address{StreetAddress: "New One"})
				gymLocation = models.GymLocation{AddressId: newAddr.AddressId, LocationName: locationName}
				created, _ = datastore.CreateGymLocation(gymLocation)
			})

			AfterEach(func() {
				datastore.DeleteGymLocation(created.GymLocationId)
				datastore.DeleteAddress(newAddr.AddressId)
			})

			It("should return the created gymLocation", func() {
				Expect(created.LocationName).To(Equal(locationName))
			})

			It("should add a gymLocation to the db", func() {
				newGymLocation, _ := datastore.GetGymLocation(created.GymLocationId)
				Expect(newGymLocation.LocationName).To(Equal(locationName))
			})
		})

		Describe("Unsuccessful call", func() {
			It("should return an error object if gymLocation is not unique", func() {
				street := "Test Street"
				loc := models.GymLocation{AddressId: addr.AddressId, LocationName: street}
				_, err := datastore.CreateGymLocation(loc)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("UpdateGymLocation", func() {
		var (
			locationName string = "Update Location"
			gymLocation  models.GymLocation
			created      *models.GymLocation
			updated      *models.GymLocation
			err          error
			newAddr      *models.Address
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				newAddr, _ = datastore.CreateAddress(models.Address{StreetAddress: "Another Test Street"})
				gymLocation = models.GymLocation{AddressId: newAddr.AddressId, LocationName: locationName}
				created, _ = datastore.CreateGymLocation(models.GymLocation{
					AddressId:    newAddr.AddressId,
					LocationName: "Test Name",
				})
				updated, _ = datastore.UpdateGymLocation(created.GymLocationId, gymLocation)
			})

			AfterEach(func() {
				datastore.DeleteGymLocation(updated.GymLocationId)
				datastore.DeleteAddress(newAddr.AddressId)
			})

			It("should return the updated gymLocation", func() {
				Expect(updated.LocationName).To(Equal(locationName))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				gymLocation = models.GymLocation{LocationName: "Test Name"}
				updated, err = datastore.UpdateGymLocation(2, gymLocation)
			})

			It("should return an error object if gymLocation to update doesn't exist", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil gymLocation", func() {
				Expect(updated).To(BeNil())
			})
		})
	})

	Describe("DeleteGymLocation", func() {
		Describe("Successful call", func() {
			It("should return nil", func() {
				err := datastore.DeleteGymLocation(one.GymLocationId)
				Expect(err).To(BeNil())
			})
		})
	})
})
