package datastore_test

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day db interactions", func() {
	var dayId int64 = 1

	Describe("GetDayList", func() {
		var days []models.Day

		Describe("Successful call", func() {
			BeforeEach(func() {
				days, _ = datastore.GetDayList("")
			})

			It("should return a list of days", func() {
				Expect(len(days)).To(Equal(7))
			})
		})
	})

	Describe("GetDay", func() {
		var day *models.Day

		Describe("Successful call", func() {
			It("should return the correct day", func() {
				day, _ = datastore.GetDay(dayId)
				Expect(day.DayId).To(Equal(dayId))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentId int64 = 5000
				err           error
			)

			BeforeEach(func() {
				day, err = datastore.GetDay(nonExistentId)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil day", func() {
				Expect(day).To(BeNil())
			})
		})
	})

	Describe("GetDayCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetDayCount("")
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(7))
			})
		})
	})

	Describe("CreateDay", func() {
		var (
			dayName string = "New Day"
			day     models.Day
			created *models.Day
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				day = models.Day{DayName: dayName}
				created, _ = datastore.CreateDay(day)
			})

			AfterEach(func() {
				datastore.DeleteDay(created.DayId)
			})

			It("should return the created day", func() {
				Expect(created.DayName).To(Equal(dayName))
			})

			It("should add a day to the db", func() {
				newDay, _ := datastore.GetDay(created.DayId)
				Expect(newDay.DayName).To(Equal(dayName))
			})
		})

		Describe("Unsuccessful call", func() {
			var created *models.Day

			AfterEach(func() {
				datastore.DeleteDay(created.DayId)
			})

			It("should return an error object if day is not unique", func() {
				name := "Test Name"
				pln := models.Day{DayName: name}
				created, _ = datastore.CreateDay(pln)
				_, err := datastore.CreateDay(pln)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("UpdateDay", func() {
		var (
			dayName string = "Anytime"
			day     models.Day
			created *models.Day
			updated *models.Day
			err     error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				day = models.Day{DayName: dayName}
				created, _ = datastore.CreateDay(models.Day{DayName: "Daily"})
				updated, _ = datastore.UpdateDay(created.DayId, day)
			})

			AfterEach(func() {
				datastore.DeleteDay(updated.DayId)
			})

			It("should return the updated day", func() {
				Expect(updated.DayName).To(Equal(dayName))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				day = models.Day{DayName: "Daily"}
				updated, err = datastore.UpdateDay(10000, day)
			})

			It("should return an error object if day to update doesn't exist", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil day", func() {
				Expect(updated).To(BeNil())
			})
		})
	})

	Describe("DeleteDay", func() {
		Describe("Successful call", func() {
			It("should return nil", func() {
				created, _ := datastore.CreateDay(models.Day{DayName: "Testing"})
				err := datastore.DeleteDay(created.DayId)
				Expect(err).To(BeNil())
			})
		})
	})
})
