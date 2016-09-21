package datastore_test

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Role db interactions", func() {
	var roleID int64 = 1

	Describe("GetRoleList", func() {
		var roles []models.Role

		Describe("Successful call", func() {
			BeforeEach(func() {
				roles, _ = datastore.GetRoleList("")
			})

			It("should return a list of roles", func() {
				Expect(len(roles)).To(Equal(5))
			})
		})
	})

	Describe("GetRole", func() {
		var role *models.Role

		Describe("Successful call", func() {
			It("should return the correct role", func() {
				role, _ = datastore.GetRole(roleID)
				Expect(role.RoleID).To(Equal(roleID))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentID int64 = 5000
				err           error
			)

			BeforeEach(func() {
				role, err = datastore.GetRole(nonExistentID)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil role", func() {
				Expect(role).To(BeNil())
			})
		})
	})

	Describe("GetRoleCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetRoleCount("")
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(5))
			})
		})
	})

	Describe("CreateRole", func() {
		var (
			roleName string = "New Role"
			role     models.Role
			created  *models.Role
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				role = models.Role{RoleName: roleName}
				created, _ = datastore.CreateRole(role)
			})

			AfterEach(func() {
				datastore.DeleteRole(created.RoleID)
			})

			It("should return the created role", func() {
				Expect(created.RoleName).To(Equal(roleName))
			})

			It("should add a role to the db", func() {
				newRole, _ := datastore.GetRole(created.RoleID)
				Expect(newRole.RoleName).To(Equal(roleName))
			})
		})

		Describe("Unsuccessful call", func() {
			var created *models.Role

			AfterEach(func() {
				datastore.DeleteRole(created.RoleID)
			})

			It("should return an error object if role is not unique", func() {
				name := "Test Name"
				pln := models.Role{RoleName: name}
				created, _ = datastore.CreateRole(pln)
				_, err := datastore.CreateRole(pln)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("UpdateRole", func() {
		var (
			roleName string = "Anytime"
			role     models.Role
			created  *models.Role
			updated  *models.Role
			err      error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				role = models.Role{RoleName: roleName}
				created, _ = datastore.CreateRole(models.Role{RoleName: "Daily"})
				updated, _ = datastore.UpdateRole(created.RoleID, role)
			})

			AfterEach(func() {
				datastore.DeleteRole(updated.RoleID)
			})

			It("should return the updated role", func() {
				Expect(updated.RoleName).To(Equal(roleName))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				role = models.Role{RoleName: "Daily"}
				updated, err = datastore.UpdateRole(10000, role)
			})

			It("should return an error object if role to update doesn't exist", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil role", func() {
				Expect(updated).To(BeNil())
			})
		})
	})

	Describe("DeleteRole", func() {
		Describe("Successful call", func() {
			It("should return nil", func() {
				created, _ := datastore.CreateRole(models.Role{RoleName: "Testing"})
				err := datastore.DeleteRole(created.RoleID)
				Expect(err).To(BeNil())
			})
		})
	})
})
