package datastore_test

import (
	"github.com/lukashambsch/anygym.api/models"
	"github.com/lukashambsch/anygym.api/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OutsideMembership db interactions", func() {
	var (
		outsideMembershipOne, outsideMembershipTwo *models.OutsideMembership
		memberID, gymOneID, gymTwoID               int64 = 1, 1, 2
	)

	BeforeEach(func() {
		outsideMembershipOne, _ = datastore.CreateOutsideMembership(models.OutsideMembership{
			GymID:    &gymOneID,
			MemberID: memberID,
		})
		outsideMembershipTwo, _ = datastore.CreateOutsideMembership(models.OutsideMembership{
			GymID:    &gymTwoID,
			MemberID: memberID,
		})
	})

	AfterEach(func() {
		datastore.DeleteOutsideMembership(outsideMembershipOne.OutsideMembershipID)
		datastore.DeleteOutsideMembership(outsideMembershipTwo.OutsideMembershipID)
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
				outsideMembership, _ = datastore.GetOutsideMembership(outsideMembershipOne.OutsideMembershipID)
				Expect(outsideMembership.OutsideMembershipID).To(Equal(outsideMembershipOne.OutsideMembershipID))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentID int64 = 5000
				err           error
			)

			BeforeEach(func() {
				outsideMembership, err = datastore.GetOutsideMembership(nonExistentID)
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
			otherGymID        int64 = 3
			outsideMembership models.OutsideMembership
			created           *models.OutsideMembership
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				outsideMembership = models.OutsideMembership{
					GymID:    &otherGymID,
					MemberID: memberID,
				}
				created, _ = datastore.CreateOutsideMembership(outsideMembership)
			})

			AfterEach(func() {
				datastore.DeleteOutsideMembership(created.OutsideMembershipID)
			})

			It("should return the created outsideMembership", func() {
				Expect(created.GymID).To(Equal(&otherGymID))
			})

			It("should add a outsideMembership to the db", func() {
				newMember, _ := datastore.GetOutsideMembership(created.OutsideMembershipID)
				Expect(newMember.GymID).To(Equal(&otherGymID))
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
			gymOneID, gymTwoID int64 = 3, 4
			outsideMembership  models.OutsideMembership
			created            *models.OutsideMembership
			updated            *models.OutsideMembership
			err                error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				created, _ = datastore.CreateOutsideMembership(models.OutsideMembership{
					GymID:    &gymOneID,
					MemberID: memberID,
				})
				created.GymID = &gymTwoID
				updated, _ = datastore.UpdateOutsideMembership(created.OutsideMembershipID, *created)
			})

			AfterEach(func() {
				datastore.DeleteOutsideMembership(updated.OutsideMembershipID)
			})

			It("should return the updated outsideMembership", func() {
				Expect(updated.GymID).To(Equal(&gymTwoID))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				outsideMembership = models.OutsideMembership{
					GymID:    &gymOneID,
					MemberID: memberID,
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
				var otherGymID int64 = 3
				created, _ := datastore.CreateOutsideMembership(models.OutsideMembership{
					GymID:    &otherGymID,
					MemberID: memberID,
				})
				err := datastore.DeleteOutsideMembership(created.OutsideMembershipID)
				Expect(err).To(BeNil())
			})
		})
	})
})
