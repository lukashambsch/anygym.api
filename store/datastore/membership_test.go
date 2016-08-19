package datastore_test

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Membership db interactions", func() {
	var membershipId, planId, memberId int64 = 1, 1, 1

	Describe("GetMembershipList", func() {
		var memberships []models.Membership

		Describe("Successful call", func() {
			BeforeEach(func() {
				memberships, _ = datastore.GetMembershipList()
			})

			It("should return a list of memberships", func() {
				Expect(len(memberships)).To(Equal(2))
			})
		})
	})

	Describe("GetMembership", func() {
		var membership *models.Membership

		Describe("Successful call", func() {
			It("should return the correct membership", func() {
				membership, _ = datastore.GetMembership(membershipId)
				Expect(membership.MembershipId).To(Equal(membershipId))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentId int64 = 5
				err           error
			)

			BeforeEach(func() {
				membership, err = datastore.GetMembership(nonExistentId)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil membership", func() {
				Expect(membership).To(BeNil())
			})
		})
	})

	Describe("GetMembershipCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetMembershipCount()
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(2))
			})
		})
	})

	Describe("CreateMembership", func() {
		var (
			active     bool = true
			membership models.Membership
			created    *models.Membership
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				membership = models.Membership{PlanId: &planId, MemberId: &memberId, Active: active}
				created, _ = datastore.CreateMembership(membership)
			})

			AfterEach(func() {
				datastore.DeleteMembership(created.MembershipId)
			})

			It("should return the created membership", func() {
				Expect(created.Active).To(Equal(active))
			})

			It("should add a membership to the db", func() {
				newMember, _ := datastore.GetMembership(created.MembershipId)
				Expect(newMember.Active).To(Equal(active))
			})
		})

		Describe("Unsuccessful call", func() {
			It("should return an error object if no plan_id and member_id", func() {
				mbr := models.Membership{}
				_, err := datastore.CreateMembership(mbr)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("UpdateMembership", func() {
		var (
			active     bool = true
			membership models.Membership
			created    *models.Membership
			updated    *models.Membership
			err        error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				created, _ = datastore.CreateMembership(
					models.Membership{PlanId: &planId, MemberId: &memberId},
				)
				created.Active = active
				updated, _ = datastore.UpdateMembership(created.MembershipId, *created)
			})

			AfterEach(func() {
				datastore.DeleteMembership(updated.MembershipId)
			})

			It("should return the updated membership", func() {
				Expect(updated.Active).To(Equal(active))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				membership = models.Membership{PlanId: &planId, MemberId: &memberId}
				updated, err = datastore.UpdateMembership(3, membership)
			})

			It("should return an error object if membership to update doesn't exist", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil membership", func() {
				Expect(updated).To(BeNil())
			})
		})
	})

	Describe("DeleteMembership", func() {
		Describe("Successful call", func() {
			It("should return nil", func() {
				err := datastore.DeleteMembership(membershipId)
				Expect(err).To(BeNil())
			})
		})
	})
})
