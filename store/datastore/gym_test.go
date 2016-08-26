package datastore_test

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gym db interactions", func() {
	var gymId int64 = 1

	Describe("GetGymList", func() {
		var gyms []models.Gym

		Describe("Successful call", func() {
			BeforeEach(func() {
				gyms, _ = datastore.GetGymList("")
			})

			It("should return a list of gyms", func() {
				Expect(len(gyms)).To(Equal(4))
			})
		})
	})

	Describe("GetGym", func() {
		var gym *models.Gym

		Describe("Successful call", func() {
			It("should return the correct gym", func() {
				gym, _ = datastore.GetGym(gymId)
				Expect(gym.GymId).To(Equal(gymId))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentId int64 = 5000
				err           error
			)

			BeforeEach(func() {
				gym, err = datastore.GetGym(nonExistentId)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil gym", func() {
				Expect(gym).To(BeNil())
			})
		})
	})

	Describe("GetGymCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetGymCount("")
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(4))
			})
		})
	})

	Describe("CreateGym", func() {
		var (
			gymName string = "New Gym"
			gym     models.Gym
			created *models.Gym
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				gym = models.Gym{GymName: gymName}
				created, _ = datastore.CreateGym(gym)
			})

			AfterEach(func() {
				datastore.DeleteGym(created.GymId)
			})

			It("should return the created gym", func() {
				Expect(created.GymName).To(Equal(gymName))
			})

			It("should add a gym to the db", func() {
				newGym, _ := datastore.GetGym(created.GymId)
				Expect(newGym.GymName).To(Equal(gymName))
			})
		})

		Describe("Unsuccessful call", func() {
		})
	})

	Describe("UpdateGym", func() {
		var (
			gymName string = "Fitness"
			gym     models.Gym
			created *models.Gym
			updated *models.Gym
			err     error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				gym = models.Gym{GymName: gymName}
				created, _ = datastore.CreateGym(models.Gym{GymName: "Gym"})
				updated, _ = datastore.UpdateGym(created.GymId, gym)
			})

			AfterEach(func() {
				datastore.DeleteGym(updated.GymId)
			})

			It("should return the updated gym", func() {
				Expect(updated.GymName).To(Equal(gymName))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				gym = models.Gym{GymName: "Pending"}
				updated, err = datastore.UpdateGym(2000, gym)
			})

			It("should return an error object", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil gym", func() {
				Expect(updated).To(BeNil())
			})
		})
	})

	Describe("DeleteGym", func() {
		Describe("Successful call", func() {
			It("should return nil", func() {
				created, _ := datastore.CreateGym(models.Gym{GymName: "Test"})
				err := datastore.DeleteGym(created.GymId)
				Expect(err).To(BeNil())
			})
		})
	})
})
