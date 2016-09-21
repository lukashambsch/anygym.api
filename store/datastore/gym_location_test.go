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
		gymID         int64 = 1
	)

	BeforeEach(func() {
		addr, _ = datastore.CreateAddress(models.Address{StreetAddress: "Testing"})
		address, _ = datastore.CreateAddress(models.Address{StreetAddress: "Testing Two"})
		one, _ = datastore.CreateGymLocation(models.GymLocation{
			GymID:        gymID,
			AddressID:    addr.AddressID,
			LocationName: "Testing",
		})
		two, _ = datastore.CreateGymLocation(models.GymLocation{
			GymID:        gymID,
			AddressID:    address.AddressID,
			LocationName: "Testing Two",
		})
	})

	AfterEach(func() {
		datastore.DeleteGymLocation(one.GymLocationID)
		datastore.DeleteGymLocation(two.GymLocationID)
		datastore.DeleteAddress(addr.AddressID)
		datastore.DeleteAddress(address.AddressID)
	})

	Describe("GetGymLocationList", func() {
		var gymLocations []models.GymLocation

		Describe("Successful call", func() {
			BeforeEach(func() {
				gymLocations, _ = datastore.GetGymLocationList("")
			})

			It("should return a list of gymLocations", func() {
				Expect(len(gymLocations)).To(Equal(4))
			})
		})
	})

	Describe("GetGymLocation", func() {
		var gymLocation *models.GymLocation

		Describe("Successful call", func() {
			It("should return the correct gymLocation", func() {
				gymLocation, _ = datastore.GetGymLocation(one.GymLocationID)
				Expect(gymLocation.GymLocationID).To(Equal(one.GymLocationID))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentID int64 = 5000
				err           error
			)

			BeforeEach(func() {
				gymLocation, err = datastore.GetGymLocation(nonExistentID)
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
				Expect(*count).To(Equal(4))
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
				gymLocation = models.GymLocation{GymID: gymID, AddressID: newAddr.AddressID, LocationName: locationName}
				created, _ = datastore.CreateGymLocation(gymLocation)
			})

			AfterEach(func() {
				datastore.DeleteGymLocation(created.GymLocationID)
				datastore.DeleteAddress(newAddr.AddressID)
			})

			It("should return the created gymLocation", func() {
				Expect(created.LocationName).To(Equal(locationName))
			})

			It("should add a gymLocation to the db", func() {
				newGymLocation, _ := datastore.GetGymLocation(created.GymLocationID)
				Expect(newGymLocation.LocationName).To(Equal(locationName))
			})
		})

		Describe("Unsuccessful call", func() {
			It("should return an error object if gymLocation is not unique", func() {
				street := "Test Street"
				loc := models.GymLocation{GymID: gymID, AddressID: addr.AddressID, LocationName: street}
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
				gymLocation = models.GymLocation{GymID: gymID, AddressID: newAddr.AddressID, LocationName: locationName}
				created, _ = datastore.CreateGymLocation(models.GymLocation{
					GymID:        gymID,
					AddressID:    newAddr.AddressID,
					LocationName: "Test Name",
				})
				updated, _ = datastore.UpdateGymLocation(created.GymLocationID, gymLocation)
			})

			AfterEach(func() {
				datastore.DeleteGymLocation(updated.GymLocationID)
				datastore.DeleteAddress(newAddr.AddressID)
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
				err := datastore.DeleteGymLocation(one.GymLocationID)
				Expect(err).To(BeNil())
			})
		})
	})
})
