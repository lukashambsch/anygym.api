package datastore_test

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Device db interactions", func() {
	var one, two, three, four *models.Device

	BeforeEach(func() {
		one, _ = datastore.CreateDevice(models.Device{UserId: 1, DeviceToken: "Testing"})
		two, _ = datastore.CreateDevice(models.Device{UserId: 1, DeviceToken: "Testing Two"})
		three, _ = datastore.CreateDevice(models.Device{UserId: 1, DeviceToken: "Testing Three"})
		four, _ = datastore.CreateDevice(models.Device{UserId: 1, DeviceToken: "Testing Four"})
	})

	AfterEach(func() {
		datastore.DeleteDevice(one.DeviceId)
		datastore.DeleteDevice(two.DeviceId)
		datastore.DeleteDevice(three.DeviceId)
		datastore.DeleteDevice(four.DeviceId)
	})

	Describe("GetDeviceList", func() {
		var devices []models.Device

		Describe("Successful call", func() {
			BeforeEach(func() {
				devices, _ = datastore.GetDeviceList("")
			})

			It("should return a list of devices", func() {
				Expect(len(devices)).To(Equal(4))
			})
		})
	})

	Describe("GetDevice", func() {
		var device *models.Device

		Describe("Successful call", func() {
			It("should return the correct device", func() {
				device, _ = datastore.GetDevice(one.DeviceId)
				Expect(device.DeviceId).To(Equal(one.DeviceId))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentId int64 = 5
				err           error
			)

			BeforeEach(func() {
				device, err = datastore.GetDevice(nonExistentId)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil device", func() {
				Expect(device).To(BeNil())
			})
		})
	})

	Describe("GetDeviceCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetDeviceCount()
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(4))
			})
		})
	})

	Describe("CreateDevice", func() {
		var (
			deviceToken string = "New Device"
			device      models.Device
			created     *models.Device
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				device = models.Device{UserId: 1, DeviceToken: deviceToken}
				created, _ = datastore.CreateDevice(device)
			})

			AfterEach(func() {
				datastore.DeleteDevice(created.DeviceId)
			})

			It("should return the created device", func() {
				Expect(created.DeviceToken).To(Equal(deviceToken))
			})

			It("should add a device to the db", func() {
				newDevice, _ := datastore.GetDevice(created.DeviceId)
				Expect(newDevice.DeviceToken).To(Equal(deviceToken))
			})
		})

		Describe("Unsuccessful call", func() {
			It("should return an error object if device is not unique", func() {
				name := "Test Name"
				pln := models.Device{DeviceToken: name}
				datastore.CreateDevice(pln)
				_, err := datastore.CreateDevice(pln)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("UpdateDevice", func() {
		var (
			deviceToken string = "Anytime"
			device      models.Device
			created     *models.Device
			updated     *models.Device
			err         error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				device = models.Device{UserId: 2, DeviceToken: deviceToken}
				created, _ = datastore.CreateDevice(models.Device{UserId: 1, DeviceToken: "Daily"})
				updated, _ = datastore.UpdateDevice(created.DeviceId, device)
			})

			AfterEach(func() {
				datastore.DeleteDevice(updated.DeviceId)
			})

			It("should return the updated device", func() {
				Expect(updated.DeviceToken).To(Equal(deviceToken))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				device = models.Device{UserId: 1, DeviceToken: "Daily"}
				updated, err = datastore.UpdateDevice(10000, device)
			})

			It("should return an error object if device to update doesn't exist", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil device", func() {
				Expect(updated).To(BeNil())
			})
		})
	})

	Describe("DeleteDevice", func() {
		Describe("Successful call", func() {
			It("should return nil", func() {
				err := datastore.DeleteDevice(one.DeviceId)
				Expect(err).To(BeNil())
			})
		})
	})
})
