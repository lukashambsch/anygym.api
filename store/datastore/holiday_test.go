package datastore_test

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Holiday db interactions", func() {
	var holidayId int64 = 1

	Describe("GetHolidayList", func() {
		var holidays []models.Holiday

		Describe("Successful call", func() {
			BeforeEach(func() {
				holidays, _ = datastore.GetHolidayList("")
			})

			It("should return a list of holidays", func() {
				Expect(len(holidays)).To(Equal(12))
			})
		})
	})

	Describe("GetHoliday", func() {
		var holiday *models.Holiday

		Describe("Successful call", func() {
			It("should return the correct holiday", func() {
				holiday, _ = datastore.GetHoliday(holidayId)
				Expect(holiday.HolidayId).To(Equal(holidayId))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentId int64 = 5000
				err           error
			)

			BeforeEach(func() {
				holiday, err = datastore.GetHoliday(nonExistentId)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil holiday", func() {
				Expect(holiday).To(BeNil())
			})
		})
	})

	Describe("GetHolidayCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetHolidayCount("")
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(12))
			})
		})
	})

	Describe("CreateHoliday", func() {
		var (
			holidayName string = "New Holiday"
			holiday     models.Holiday
			created     *models.Holiday
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				holiday = models.Holiday{HolidayName: holidayName}
				created, _ = datastore.CreateHoliday(holiday)
			})

			AfterEach(func() {
				datastore.DeleteHoliday(created.HolidayId)
			})

			It("should return the created holiday", func() {
				Expect(created.HolidayName).To(Equal(holidayName))
			})

			It("should add a holiday to the db", func() {
				newHoliday, _ := datastore.GetHoliday(created.HolidayId)
				Expect(newHoliday.HolidayName).To(Equal(holidayName))
			})
		})

		Describe("Unsuccessful call", func() {
			var created *models.Holiday

			AfterEach(func() {
				datastore.DeleteHoliday(created.HolidayId)
			})

			It("should return an error object if holiday is not unique", func() {
				name := "Test Name"
				pln := models.Holiday{HolidayName: name}
				created, _ = datastore.CreateHoliday(pln)
				_, err := datastore.CreateHoliday(pln)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("UpdateHoliday", func() {
		var (
			holidayName string = "Anytime"
			holiday     models.Holiday
			created     *models.Holiday
			updated     *models.Holiday
			err         error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				holiday = models.Holiday{HolidayName: holidayName}
				created, _ = datastore.CreateHoliday(models.Holiday{HolidayName: "Daily"})
				updated, _ = datastore.UpdateHoliday(created.HolidayId, holiday)
			})

			AfterEach(func() {
				datastore.DeleteHoliday(updated.HolidayId)
			})

			It("should return the updated holiday", func() {
				Expect(updated.HolidayName).To(Equal(holidayName))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				holiday = models.Holiday{HolidayName: "Daily"}
				updated, err = datastore.UpdateHoliday(10000, holiday)
			})

			It("should return an error object if holiday to update doesn't exist", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil holiday", func() {
				Expect(updated).To(BeNil())
			})
		})
	})

	Describe("DeleteHoliday", func() {
		Describe("Successful call", func() {
			It("should return nil", func() {
				created, _ := datastore.CreateHoliday(models.Holiday{HolidayName: "Testing"})
				err := datastore.DeleteHoliday(created.HolidayId)
				Expect(err).To(BeNil())
			})
		})
	})
})
