package datastore_test

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OutsideMembership db interactions", func() {
	var (
		outsideMembershipOne, outsideMembershipTwo *models.OutsideMembership
		memberId, gymOneId, gymTwoId               int64 = 1, 1, 2
	)

	BeforeEach(func() {
		outsideMembershipOne, _ = datastore.CreateOutsideMembership(models.OutsideMembership{
			GymId:    &gymOneId,
			MemberId: memberId,
		})
		outsideMembershipTwo, _ = datastore.CreateOutsideMembership(models.OutsideMembership{
			GymId:    &gymTwoId,
			MemberId: memberId,
		})
	})

	AfterEach(func() {
		datastore.DeleteOutsideMembership(outsideMembershipOne.OutsideMembershipId)
		datastore.DeleteOutsideMembership(outsideMembershipTwo.OutsideMembershipId)
	})

	Describe("GetOutsideMembershipList", func() {
		var outsideMemberships []models.OutsideMembership

		Describe("Successful call", func() {
			BeforeEach(func() {
				outsideMemberships, _ = datastore.GetOutsideMembershipList("")
			})

			It("should return a list of outsideMemberships", func() {
				Expect(len(outsideMemberships)).To(Equal(2))
			})
		})
	})

	Describe("GetOutsideMembership", func() {
		var outsideMembership *models.OutsideMembership

		Describe("Successful call", func() {
			It("should return the correct outsideMembership", func() {
				outsideMembership, _ = datastore.GetOutsideMembership(outsideMembershipOne.OutsideMembershipId)
				Expect(outsideMembership.OutsideMembershipId).To(Equal(outsideMembershipOne.OutsideMembershipId))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentId int64 = 5000
				err           error
			)

			BeforeEach(func() {
				outsideMembership, err = datastore.GetOutsideMembership(nonExistentId)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil outsideMembership", func() {
				Expect(outsideMembership).To(BeNil())
			})
		})
	})

	Describe("GetOutsideMembershipCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetOutsideMembershipCount("")
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(2))
			})
		})
	})

	Describe("CreateOutsideMembership", func() {
		var (
			otherGymId        int64 = 3
			outsideMembership models.OutsideMembership
			created           *models.OutsideMembership
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				outsideMembership = models.OutsideMembership{
					GymId:    &otherGymId,
					MemberId: memberId,
				}
				created, _ = datastore.CreateOutsideMembership(outsideMembership)
			})

			AfterEach(func() {
				datastore.DeleteOutsideMembership(created.OutsideMembershipId)
			})

			It("should return the created outsideMembership", func() {
				Expect(created.GymId).To(Equal(&otherGymId))
			})

			It("should add a outsideMembership to the db", func() {
				newMember, _ := datastore.GetOutsideMembership(created.OutsideMembershipId)
				Expect(newMember.GymId).To(Equal(&otherGymId))
			})
		})

		Describe("Unsuccessful call", func() {
			It("should return an error object if no day_id or holiday_id", func() {
				mbr := models.OutsideMembership{}
				_, err := datastore.CreateOutsideMembership(mbr)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("UpdateOutsideMembership", func() {
		var (
			gymOneId, gymTwoId int64 = 3, 4
			outsideMembership  models.OutsideMembership
			created            *models.OutsideMembership
			updated            *models.OutsideMembership
			err                error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				created, _ = datastore.CreateOutsideMembership(models.OutsideMembership{
					GymId:    &gymOneId,
					MemberId: memberId,
				})
				created.GymId = &gymTwoId
				updated, _ = datastore.UpdateOutsideMembership(created.OutsideMembershipId, *created)
			})

			AfterEach(func() {
				datastore.DeleteOutsideMembership(updated.OutsideMembershipId)
			})

			It("should return the updated outsideMembership", func() {
				Expect(updated.GymId).To(Equal(&gymTwoId))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				outsideMembership = models.OutsideMembership{
					GymId:    &gymOneId,
					MemberId: memberId,
				}
				updated, err = datastore.UpdateOutsideMembership(5000, outsideMembership)
			})

			It("should return an error object if outsideMembership to update doesn't exist", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil outsideMembership", func() {
				Expect(updated).To(BeNil())
			})
		})
	})

	Describe("DeleteOutsideMembership", func() {
		Describe("Successful call", func() {
			It("should return nil", func() {
				var otherGymId int64 = 3
				created, _ := datastore.CreateOutsideMembership(models.OutsideMembership{
					GymId:    &otherGymId,
					MemberId: memberId,
				})
				err := datastore.DeleteOutsideMembership(created.OutsideMembershipId)
				Expect(err).To(BeNil())
			})
		})
	})
})
