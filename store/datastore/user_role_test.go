package datastore_test

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserRole db interactions", func() {
	var userRoleId int64 = 1

	Describe("GetUserRoleList", func() {
		var userRoles []models.UserRole

		Describe("Successful call", func() {
			BeforeEach(func() {
				userRoles, _ = datastore.GetUserRoleList("")
			})

			It("should return a list of userRoles", func() {
				Expect(len(userRoles)).To(Equal(2))
			})
		})
	})

	Describe("GetUserRole", func() {
		var userRole *models.UserRole

		Describe("Successful call", func() {
			It("should return the correct userRole", func() {
				userRole, _ = datastore.GetUserRole(userRoleId)
				Expect(userRole.UserRoleId).To(Equal(userRoleId))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentId int64 = 5
				err           error
			)

			BeforeEach(func() {
				userRole, err = datastore.GetUserRole(nonExistentId)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil plan", func() {
				Expect(userRole).To(BeNil())
			})
		})
	})

	Describe("GetUserRoleCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetUserRoleCount()
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(2))
			})
		})
	})

	Describe("CreateUserRole", func() {
		var (
			userId   int64 = 1
			userRole models.UserRole
			created  *models.UserRole
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				userRole = models.UserRole{UserId: userId, RoleId: 5}
				created, _ = datastore.CreateUserRole(userRole)
			})

			AfterEach(func() {
				datastore.DeleteUserRole(created.UserRoleId)
			})

			It("should return the created userRole", func() {
				Expect(created.UserId).To(Equal(userId))
			})

			It("should add a userRole to the db", func() {
				newUserRole, _ := datastore.GetUserRole(created.UserRoleId)
				Expect(newUserRole.UserId).To(Equal(userId))
			})
		})

		Describe("Unsuccessful call", func() {
			It("should return an error object if userRole is not unique", func() {
				usrRole := models.UserRole{UserId: userId, RoleId: 1}
				datastore.CreateUserRole(usrRole)
				_, err := datastore.CreateUserRole(usrRole)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("UpdateUserRole", func() {
		var (
			userId   int64 = 2
			userRole models.UserRole
			created  *models.UserRole
			updated  *models.UserRole
			err      error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				userRole = models.UserRole{UserId: userId, RoleId: 5}
				created, _ = datastore.CreateUserRole(models.UserRole{UserId: 1, RoleId: 5})
				updated, _ = datastore.UpdateUserRole(created.UserRoleId, userRole)
			})

			AfterEach(func() {
				datastore.DeleteUserRole(updated.UserRoleId)
			})

			It("should return the updated userRole", func() {
				Expect(updated.UserId).To(Equal(userId))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				userRole = models.UserRole{}
				updated, err = datastore.UpdateUserRole(10000, userRole)
			})

			It("should return an error object if userRole to update doesn't exist", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil userRole", func() {
				Expect(updated).To(BeNil())
			})
		})
	})

	Describe("DeleteUserRole", func() {
		Describe("Successful call", func() {
			It("should return nil", func() {
				created, _ := datastore.CreateUserRole(models.UserRole{UserId: 1, RoleId: 5})
				err := datastore.DeleteUserRole(created.UserRoleId)
				Expect(err).To(BeNil())
			})
		})
	})
})
