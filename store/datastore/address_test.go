package datastore_test

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
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
		datastore.DeleteAddress(one.AddressId)
		datastore.DeleteAddress(two.AddressId)
		datastore.DeleteAddress(three.AddressId)
		datastore.DeleteAddress(four.AddressId)
	})

	Describe("GetAddressList", func() {
		var addresses []models.Address

		Describe("Successful call", func() {
			BeforeEach(func() {
				addresses, _ = datastore.GetAddressList("")
			})

			It("should return a list of addresses", func() {
				Expect(len(addresses)).To(Equal(4))
			})
		})
	})

	Describe("GetAddress", func() {
		var address *models.Address

		Describe("Successful call", func() {
			It("should return the correct address", func() {
				address, _ = datastore.GetAddress(one.AddressId)
				Expect(address.AddressId).To(Equal(one.AddressId))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentId int64 = 5
				err           error
			)

			BeforeEach(func() {
				address, err = datastore.GetAddress(nonExistentId)
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
				count, _ = datastore.GetAddressCount()
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(4))
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
				datastore.DeleteAddress(created.AddressId)
			})

			It("should return the created address", func() {
				Expect(created.StreetAddress).To(Equal(streetAddress))
			})

			It("should add a address to the db", func() {
				newAddress, _ := datastore.GetAddress(created.AddressId)
				Expect(newAddress.StreetAddress).To(Equal(streetAddress))
			})
		})

		Describe("Unsuccessful call", func() {
			It("should return an error object if address is not unique", func() {
				street := "Test Street"
				addr := models.Address{StreetAddress: street}
				datastore.CreateAddress(addr)
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
				updated, _ = datastore.UpdateAddress(created.AddressId, address)
			})

			AfterEach(func() {
				datastore.DeleteAddress(updated.AddressId)
			})

			It("should return the updated address", func() {
				Expect(updated.StreetAddress).To(Equal(streetAddress))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				address = models.Address{StreetAddress: "456 Test Ave."}
				updated, err = datastore.UpdateAddress(2, address)
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
				err := datastore.DeleteAddress(one.AddressId)
				Expect(err).To(BeNil())
			})
		})
	})
})
