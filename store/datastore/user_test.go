package datastore_test

import (
	"github.com/lukashambsch/anygym.api/models"
	"github.com/lukashambsch/anygym.api/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User db interactions", func() {
	var userID int64 = 1

	Describe("GetUserList", func() {
		var users []models.User

		Describe("Successful call", func() {
			BeforeEach(func() {
				users, _ = datastore.GetUserList("")
			})

			It("should return a list of users", func() {
				Expect(len(users)).To(Equal(2))
			})
		})
	})

	Describe("GetUser", func() {
		var user *models.User

		Describe("Successful call", func() {
			It("should return the correct user", func() {
				user, _ = datastore.GetUser(userID)
				Expect(user.UserID).To(Equal(userID))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentID int64 = 5
				err           error
			)

			BeforeEach(func() {
				user, err = datastore.GetUser(nonExistentID)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil user", func() {
				Expect(user).To(BeNil())
			})
		})
	})

	Describe("GetUserRoles", func() {
		var user *models.User

		Describe("Successful call", func() {
			It("should return the list of the user's roles", func() {
				user, _ = datastore.GetUser(userID)
				roles, _ := datastore.GetUserRoles(user.UserID)
				Expect(len(roles)).To(Equal(1))
			})
		})
	})

	Describe("GetUserCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetUserCount("")
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(2))
			})
		})
	})

	Describe("CreateUser", func() {
		var (
			email   string = "test@gmail.com"
			user    models.User
			created *models.User
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				user = models.User{Email: email}
				created, _ = datastore.CreateUser(user)
			})

			AfterEach(func() {
				datastore.DeleteUser(created.UserID)
			})

			It("should return the created user", func() {
				Expect(created.Email).To(Equal(email))
			})

			It("should add a user to the db", func() {
				newMember, _ := datastore.GetUser(created.UserID)
				Expect(newMember.Email).To(Equal(email))
			})
		})

		Describe("Unsuccessful call", func() {
			It("should return an error object if no email", func() {
				mbr := models.User{}
				_, err := datastore.CreateUser(mbr)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("UpdateUser", func() {
		var (
			email   string = "test@gmail.com"
			user    models.User
			created *models.User
			updated *models.User
			err     error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				created, _ = datastore.CreateUser(models.User{Email: "different"})
				created.Email = email
				updated, _ = datastore.UpdateUser(created.UserID, *created)
			})

			AfterEach(func() {
				datastore.DeleteUser(updated.UserID)
			})

			It("should return the updated user", func() {
				Expect(updated.Email).To(Equal(email))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				user = models.User{Email: email}
				updated, err = datastore.UpdateUser(3, user)
			})

			It("should return an error object if user to update doesn't exist", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil user", func() {
				Expect(updated).To(BeNil())
			})
		})
	})

	Describe("DeleteUser", func() {
		Describe("Successful call", func() {
			It("should return nil", func() {
				created, _ := datastore.CreateUser(models.User{Email: "email"})
				err := datastore.DeleteUser(created.UserID)
				Expect(err).To(BeNil())
			})
		})
	})
})
