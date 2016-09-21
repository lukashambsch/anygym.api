package datastore_test

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GymFeature db interactions", func() {
	var (
		one, two *models.GymFeature
		gym      *models.Gym
	)

	BeforeEach(func() {
		gym, _ = datastore.CreateGym(models.Gym{GymName: "Test Gym Name"})
		one, _ = datastore.CreateGymFeature(models.GymFeature{GymID: gym.GymID, FeatureID: 1})
		two, _ = datastore.CreateGymFeature(models.GymFeature{GymID: gym.GymID, FeatureID: 2})
	})

	AfterEach(func() {
		datastore.DeleteGymFeature(one.GymFeatureID)
		datastore.DeleteGymFeature(two.GymFeatureID)
		datastore.DeleteGym(gym.GymID)
	})

	Describe("GetGymFeatureList", func() {
		var gymFeatures []models.GymFeature

		Describe("Successful call", func() {
			BeforeEach(func() {
				gymFeatures, _ = datastore.GetGymFeatureList("")
			})

			It("should return a list of gymFeatures", func() {
				Expect(len(gymFeatures)).To(Equal(2))
			})
		})
	})

	Describe("GetGymFeature", func() {
		var gymFeature *models.GymFeature

		Describe("Successful call", func() {
			It("should return the correct gymFeature", func() {
				gymFeature, _ = datastore.GetGymFeature(one.GymFeatureID)
				Expect(gymFeature.GymFeatureID).To(Equal(one.GymFeatureID))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentID int64 = 5000
				err           error
			)

			BeforeEach(func() {
				gymFeature, err = datastore.GetGymFeature(nonExistentID)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil plan", func() {
				Expect(gymFeature).To(BeNil())
			})
		})
	})

	Describe("GetGymFeatureCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetGymFeatureCount("")
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(2))
			})
		})
	})

	Describe("CreateGymFeature", func() {
		var (
			featureID  int64 = 3
			gymFeature models.GymFeature
			created    *models.GymFeature
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				gymFeature = models.GymFeature{GymID: gym.GymID, FeatureID: featureID}
				created, _ = datastore.CreateGymFeature(gymFeature)
			})

			AfterEach(func() {
				datastore.DeleteGymFeature(created.GymFeatureID)
			})

			It("should return the created gymFeature", func() {
				Expect(created.FeatureID).To(Equal(featureID))
			})

			It("should add a gymFeature to the db", func() {
				newGymFeature, _ := datastore.GetGymFeature(created.GymFeatureID)
				Expect(newGymFeature.FeatureID).To(Equal(featureID))
			})
		})

		Describe("Unsuccessful call", func() {
			var created *models.GymFeature

			AfterEach(func() {
				datastore.DeleteGymFeature(created.GymFeatureID)
			})

			It("should return an error object if gymFeature is not unique", func() {
				gymFtr := models.GymFeature{GymID: gym.GymID, FeatureID: featureID}
				created, _ = datastore.CreateGymFeature(gymFtr)
				_, err := datastore.CreateGymFeature(gymFtr)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("UpdateGymFeature", func() {
		var (
			featureID  int64 = 18
			gymFeature models.GymFeature
			created    *models.GymFeature
			updated    *models.GymFeature
			err        error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				gymFeature = models.GymFeature{GymID: gym.GymID, FeatureID: featureID}
				created, _ = datastore.CreateGymFeature(models.GymFeature{GymID: gym.GymID, FeatureID: featureID})
				updated, _ = datastore.UpdateGymFeature(created.GymFeatureID, gymFeature)
			})

			AfterEach(func() {
				datastore.DeleteGymFeature(updated.GymFeatureID)
			})

			It("should return the updated gymFeature", func() {
				Expect(updated.FeatureID).To(Equal(featureID))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				gymFeature = models.GymFeature{}
				updated, err = datastore.UpdateGymFeature(10000, gymFeature)
			})

			It("should return an error object if gymFeature to update doesn't exist", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil gymFeature", func() {
				Expect(updated).To(BeNil())
			})
		})
	})

	Describe("DeleteGymFeature", func() {
		Describe("Successful call", func() {
			It("should return nil", func() {
				created, _ := datastore.CreateGymFeature(models.GymFeature{GymID: gym.GymID, FeatureID: 10})
				err := datastore.DeleteGymFeature(created.GymFeatureID)
				Expect(err).To(BeNil())
			})
		})
	})
})
