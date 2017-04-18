package datastore_test

import (
	"github.com/lukashambsch/anygym.api/models"
	"github.com/lukashambsch/anygym.api/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Member db interactions", func() {
	var (
		memberID int64 = 1
		member   *models.Member
		user     *models.User
	)

	BeforeEach(func() {
		user, _ = datastore.CreateUser(models.User{Email: "testemail@gmail.com"})
		member, _ = datastore.CreateMember(models.Member{FirstName: "Test First", UserID: user.UserID})
	})

	AfterEach(func() {
		datastore.DeleteMember(member.MemberID)
		datastore.DeleteUser(user.UserID)
	})

	Describe("GetMemberList", func() {
		var members []models.Member

		Describe("Successful call", func() {
			BeforeEach(func() {
				members, _ = datastore.GetMemberList("")
			})

			It("should return a list of members", func() {
				Expect(len(members)).To(Equal(3))
			})
		})
	})

	Describe("GetMember", func() {
		Describe("Successful call", func() {
			It("should return the correct member", func() {
				mbr, _ := datastore.GetMember(memberID)
				Expect(mbr.MemberID).To(Equal(memberID))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentID int64 = 99999
				err           error
				mbr           *models.Member
			)

			BeforeEach(func() {
				mbr, err = datastore.GetMember(nonExistentID)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil member", func() {
				Expect(mbr).To(BeNil())
			})
		})
	})

	Describe("GetMemberCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetMemberCount("")
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(3))
			})
		})
	})

	Describe("CreateMember", func() {
		var (
			firstName string = "New Member"
			created   *models.Member
			usr       *models.User
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				usr, _ = datastore.CreateUser(models.User{Email: "Another Test"})
				mbr := models.Member{UserID: usr.UserID, FirstName: firstName}
				created, _ = datastore.CreateMember(mbr)
			})

			AfterEach(func() {
				datastore.DeleteMember(created.MemberID)
				datastore.DeleteUser(usr.UserID)
			})

			It("should return the created member", func() {
				Expect(created.FirstName).To(Equal(firstName))
			})

			It("should add a member to the db", func() {
				newMember, _ := datastore.GetMember(created.MemberID)
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
			mbr       models.Member
			usr       *models.User
			created   *models.Member
			updated   *models.Member
			err       error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				usr, _ = datastore.CreateUser(models.User{Email: "Email"})
				created, _ = datastore.CreateMember(models.Member{FirstName: "First Name", UserID: usr.UserID})
				created.FirstName = firstName
				updated, err = datastore.UpdateMember(created.MemberID, *created)
			})

			AfterEach(func() {
				datastore.DeleteMember(updated.MemberID)
				datastore.DeleteUser(usr.UserID)
			})

			It("should return the updated member", func() {
				Expect(updated.FirstName).To(Equal(firstName))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				mbr = models.Member{FirstName: "First Name"}
				updated, err = datastore.UpdateMember(2, mbr)
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
				err := datastore.DeleteMember(member.MemberID)
				Expect(err).To(BeNil())
			})
		})
	})
})
