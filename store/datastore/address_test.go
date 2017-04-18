package datastore_test

import (
	"github.com/lukashambsch/anygym.api/models"
	"github.com/lukashambsch/anygym.api/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Address db interactions", func() {
	var one, two, three, four *models.Address

	BeforeEach(func() {
		one, _ = datastore.CreateAddress(models.Address{StreetAddress: "Testing"})
		two, _ = datastore.CreateAddress(models.Address{StreetAddress: "Testing Two"})
		three, _ = datastore.CreateAddress(models.Address{StreetAddress: "Testing Three"})
		four, _ = datastore.CreateAddress(models.Address{StreetAddress: "Testing Four"})
	})

	AfterEach(func() {
		datastore.DeleteAddress(one.AddressID)
		datastore.DeleteAddress(two.AddressID)
		datastore.DeleteAddress(three.AddressID)
		datastore.DeleteAddress(four.AddressID)
	})

	Describe("GetAddressList", func() {
		var addresses []models.Address

		Describe("Successful call", func() {
			BeforeEach(func() {
				addresses, _ = datastore.GetAddressList("")
			})

			It("should return a list of addresses", func() {
				Expect(len(addresses)).To(Equal(6))
			})
		})
	})

	Describe("GetAddress", func() {
		var address *models.Address

		Describe("Successful call", func() {
			It("should return the correct address", func() {
				address, _ = datastore.GetAddress(one.AddressID)
				Expect(address.AddressID).To(Equal(one.AddressID))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentID int64 = 5
				err           error
			)

			BeforeEach(func() {
				address, err = datastore.GetAddress(nonExistentID)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil address", func() {
				Expect(address).To(BeNil())
			})
		})
	})

	Describe("GetAddressCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetAddressCount("")
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(6))
			})
		})
	})

	Describe("CreateAddress", func() {
		var (
			streetAddress string = "New Address"
			address       models.Address
			created       *models.Address
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				address = models.Address{StreetAddress: streetAddress}
				created, _ = datastore.CreateAddress(address)
			})

			AfterEach(func() {
				datastore.DeleteAddress(created.AddressID)
			})

			It("should return the created address", func() {
				Expect(created.StreetAddress).To(Equal(streetAddress))
			})

			It("should add a address to the db", func() {
				newAddress, _ := datastore.GetAddress(created.AddressID)
				Expect(newAddress.StreetAddress).To(Equal(streetAddress))
			})
		})

		Describe("Unsuccessful call", func() {
			var created *models.Address

			AfterEach(func() {
				datastore.DeleteAddress(created.AddressID)
			})

			It("should return an error object if address is not unique", func() {
				street := "Test Street"
				addr := models.Address{StreetAddress: street}
				created, _ = datastore.CreateAddress(addr)
				_, err := datastore.CreateAddress(addr)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("UpdateAddress", func() {
		var (
			streetAddress string = "123 Home St."
			address       models.Address
			created       *models.Address
			updated       *models.Address
			err           error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				address = models.Address{StreetAddress: streetAddress}
				created, _ = datastore.CreateAddress(models.Address{StreetAddress: "456 Test Ave."})
				updated, _ = datastore.UpdateAddress(created.AddressID, address)
			})

			AfterEach(func() {
				datastore.DeleteAddress(updated.AddressID)
			})

			It("should return the updated address", func() {
				Expect(updated.StreetAddress).To(Equal(streetAddress))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				address = models.Address{StreetAddress: "456 Test Ave."}
				updated, err = datastore.UpdateAddress(5000, address)
			})

			It("should return an error object if address to update doesn't exist", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil address", func() {
				Expect(updated).To(BeNil())
			})
		})
	})

	Describe("DeleteAddress", func() {
		Describe("Successful call", func() {
			It("should return nil", func() {
				err := datastore.DeleteAddress(one.AddressID)
				Expect(err).To(BeNil())
			})
		})
	})
})
