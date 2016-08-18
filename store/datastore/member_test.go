package datastore_test

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Member db interactions", func() {
	var one, two, three, four *models.Member

	BeforeEach(func() {
		one, _ = datastore.CreateMember(models.Member{FirstName: "Testing"})
		two, _ = datastore.CreateMember(models.Member{FirstName: "Testing", LastName: "Two"})
		three, _ = datastore.CreateMember(models.Member{FirstName: "Testing", LastName: "Three"})
		four, _ = datastore.CreateMember(models.Member{FirstName: "Testing", LastName: "Four"})
	})

	AfterEach(func() {
		datastore.DeleteMember(one.MemberId)
		datastore.DeleteMember(two.MemberId)
		datastore.DeleteMember(three.MemberId)
		datastore.DeleteMember(four.MemberId)
	})

	Describe("GetMemberList", func() {
		var members []models.Member

		Describe("Successful call", func() {
			BeforeEach(func() {
				members, _ = datastore.GetMemberList()
			})

			It("should return a list of members", func() {
				Expect(len(members)).To(Equal(4))
			})
		})
	})

	Describe("GetMember", func() {
		var member *models.Member

		Describe("Successful call", func() {
			It("should return the correct member", func() {
				member, _ = datastore.GetMember(one.MemberId)
				Expect(member.MemberId).To(Equal(one.MemberId))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentId int64 = 5
				err           error
			)

			BeforeEach(func() {
				member, err = datastore.GetMember(nonExistentId)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil member", func() {
				Expect(member).To(BeNil())
			})
		})
	})

	Describe("GetMemberCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetMemberCount()
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(4))
			})
		})
	})

	Describe("CreateMember", func() {
		var (
			firstName string = "New Member"
			member    models.Member
			created   *models.Member
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				member = models.Member{FirstName: firstName}
				created, _ = datastore.CreateMember(member)
			})

			AfterEach(func() {
				datastore.DeleteMember(created.MemberId)
			})

			It("should return the created member", func() {
				Expect(created.FirstName).To(Equal(firstName))
			})

			It("should add a member to the db", func() {
				newMember, _ := datastore.GetMember(created.MemberId)
				Expect(newMember.FirstName).To(Equal(firstName))
			})
		})

		Describe("Unsuccessful call", func() {
			It("should return an error object if no names", func() {
				mbr := models.Member{}
				_, err := datastore.CreateMember(mbr)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("UpdateMember", func() {
		var (
			firstName string = "Test Name"
			member    models.Member
			created   *models.Member
			updated   *models.Member
			err       error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				created, _ = datastore.CreateMember(models.Member{FirstName: "First Name"})
				created.FirstName = firstName
				updated, _ = datastore.UpdateMember(created.MemberId, *created)
			})

			AfterEach(func() {
				datastore.DeleteMember(updated.MemberId)
			})

			It("should return the updated member", func() {
				Expect(updated.FirstName).To(Equal(firstName))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				member = models.Member{FirstName: "First Name"}
				updated, err = datastore.UpdateMember(2, member)
			})

			It("should return an error object if member to update doesn't exist", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil member", func() {
				Expect(updated).To(BeNil())
			})
		})
	})

	Describe("DeleteMember", func() {
		Describe("Successful call", func() {
			It("should return nil", func() {
				err := datastore.DeleteMember(one.MemberId)
				Expect(err).To(BeNil())
			})
		})
	})
})
