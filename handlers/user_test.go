package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/lukashambsch/gym-all-over/handlers"
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/router"
	"github.com/lukashambsch/gym-all-over/store"
	"github.com/lukashambsch/gym-all-over/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User API", func() {
	var (
		server     *httptest.Server
		userURL    string
		res        *http.Response
		data       []byte
		badPayload []byte = []byte(`{"user_id", 1}`)
		token      string
	)

	BeforeEach(func() {
		server = httptest.NewServer(router.Load())
		token, _ = RequestToken(server.URL)
		userURL = fmt.Sprintf("%s%s/users", server.URL, router.V1URLBase)
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("GetUsers endpoint", func() {
		var users []models.User

		Describe("Successful GET w/o query params", func() {
			BeforeEach(func() {
				res, data, _ = Request("GET", userURL, token, nil)
				json.Unmarshal(data, &users)
			})

			It("should return status code 200", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})

			It("should contain the users", func() {
				Expect(len(users)).To(Equal(2))
			})
		})

		Describe("Successful GET w/ query params", func() {
			It("should return a list of matching users - user_id", func() {
				res, data, _ = Request("GET", fmt.Sprintf("%s?user_id=1", userURL), token, nil)
				json.Unmarshal(data, &users)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(len(users)).To(Equal(1))
			})

			It("should return a matching user - user_id", func() {
				res, data, _ = Request("GET", fmt.Sprintf("%s?user_id=2", userURL), token, nil)
				json.Unmarshal(data, &users)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(len(users)).To(Equal(1))
			})

			It("should return no users with a valid field but no matches", func() {
				res, data, _ = Request("GET", fmt.Sprintf("%s?user_id=10", userURL), token, nil)
				json.Unmarshal(data, &users)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(len(users)).To(Equal(0))
			})

			It("should sort users by the correct field ascending", func() {
				res, data, _ = Request("GET", fmt.Sprintf("%s?sort_order=asc&order_by=user_id", userURL), token, nil)
				json.Unmarshal(data, &users)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(users[0].UserID).To(Equal(int64(1)))
				Expect(users[1].UserID).To(Equal(int64(2)))
			})

			It("should sort users by the correct field descending", func() {
				res, data, _ = Request("GET", fmt.Sprintf("%s?sort_order=desc&order_by=user_id", userURL), token, nil)
				json.Unmarshal(data, &users)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(users[0].UserID).To(Equal(int64(2)))
				Expect(users[1].UserID).To(Equal(int64(1)))
			})
		})

		Describe("Unsuccessful GET w/ query params", func() {
			var errRes handlers.APIErrorMessage

			It("should return an error with an invalid field as query param", func() {
				res, data, _ = Request("GET", fmt.Sprintf("%s?invalid=test", userURL), token, nil)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusNotFound))
				Expect(errRes.Message).To(Equal("Invalid field in query params."))
			})

			It("should return an error with an invalid field in order_by", func() {
				res, data, _ = Request("GET", fmt.Sprintf("%s?order_by=invalid", userURL), token, nil)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusNotFound))
				Expect(errRes.Message).To(Equal("Invalid order_by field."))
			})

			It("should return an error with an invalid value for sort_order", func() {
				res, data, _ = Request("GET", fmt.Sprintf("%s?order_by=user_id&sort_order=random", userURL), token, nil)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusNotFound))
				Expect(errRes.Message).To(Equal("sort_order must be either 'asc', 'desc', or ''"))
			})
		})
	})

	Describe("GetUser endpoint", func() {
		var (
			user   models.User
			userID int64 = 1
		)

		Describe("Successful GET", func() {
			BeforeEach(func() {
				res, data, _ = Request("GET", fmt.Sprintf("%s/%d", userURL, userID), token, nil)
				json.Unmarshal(data, &user)
			})

			It("should return status code 200", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})

			It("should contain the user in the response", func() {
				Expect(user.UserID).To(Equal(userID))
			})
		})

		Describe("Unsuccessful GET", func() {
			var errRes handlers.APIErrorMessage

			Context("Invalid user_id", func() {
				It("should return status code 400 with a message", func() {
					res, data, _ = Request("GET", fmt.Sprintf("%s/asdf", userURL), token, nil)
					json.Unmarshal(data, &errRes)
					Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
					Expect(errRes.Message).To(Equal(handlers.InvalidUserID))
				})
			})

			Context("Non existent user_id", func() {
				It("should return status code 404 with a message", func() {
					res, data, _ = Request("GET", fmt.Sprintf("%s/10", userURL), token, nil)
					json.Unmarshal(data, &errRes)
					Expect(res.StatusCode).To(Equal(http.StatusNotFound))
					Expect(errRes.Message).ToNot(BeEmpty())
				})
			})
		})

	})

	Describe("PostUser endpoint", func() {
		var (
			user    *models.User
			payload []byte
		)

		Describe("Successful POST", func() {
			BeforeEach(func() {
				payload = []byte(fmt.Sprintf(`{"email": "new@email.com"}`))
				res, data, _ = Request("POST", userURL, token, payload)
				json.Unmarshal(data, &user)
			})

			AfterEach(func() {
				datastore.DeleteUser(user.UserID)
			})

			It("should return status code 201", func() {
				Expect(res.StatusCode).To(Equal(http.StatusCreated))
			})

			It("should contain the user", func() {
				Expect(user.Email).To(Equal("new@email.com"))
			})

			It("should save the user", func() {
				Expect(user.UserID).ToNot(Equal(int64(0)))
			})
		})

		Describe("Unsuccessful POST", func() {
			var errRes handlers.APIErrorMessage

			Describe("Bad Request", func() {
				It("should return status code 400 with a message", func() {
					res, data, _ = Request("POST", userURL, token, badPayload)
					json.Unmarshal(data, &errRes)
					Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
					Expect(errRes.Message).ToNot(BeEmpty())
				})
			})

			Describe("Internal Server Error", func() {
				It("should return status code 500 with a message", func() {
					payload = []byte(`{"user_id": 1}`)
					res, data, _ = Request("POST", userURL, token, payload)
					json.Unmarshal(data, &errRes)
					Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
					Expect(errRes.Message).ToNot(BeEmpty())
				})
			})
		})
	})

	Describe("PutUser endpoint", func() {
		var (
			user    models.User
			payload []byte = []byte(`{"user_id": 2, "email": "updated@email.com"}`)
			userID  int64  = 2
		)

		Describe("Successful PUT", func() {
			BeforeEach(func() {
				res, data, _ = Request("PUT", fmt.Sprintf("%s/%d", userURL, userID), token, payload)
				json.Unmarshal(data, &user)
			})

			AfterEach(func() {
				datastore.UpdateUser(userID, models.User{Email: "bugentry@hotmail.com.com"})
			})

			It("should return status code 200", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})

			It("should save the updated user", func() {
				updated, _ := datastore.GetUser(userID)
				Expect(updated.Email).To(Equal("updated@email.com"))
			})
		})

		Describe("Unsuccessful PUT", func() {
			var errRes handlers.APIErrorMessage

			It("should return status code 400 with a message", func() {
				res, data, _ = Request("PUT", fmt.Sprintf("%s/%d", userURL, userID), token, badPayload)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(errRes.Message).ToNot(BeEmpty())
			})

			It("should return status code 400 with a message", func() {
				res, data, _ = Request("PUT", fmt.Sprintf("%s/a", userURL), token, payload)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(errRes.Message).To(Equal(handlers.InvalidUserID))
			})

			It("should return status code 500 with a message", func() {
				res, data, _ = Request("PUT", fmt.Sprintf("%s/5000", userURL), token, payload)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
				Expect(errRes.Message).ToNot(BeEmpty())
			})
		})
	})

	Describe("DeleteUser endpoint", func() {
		var userID int64 = 111

		Describe("Successful DELETE", func() {
			BeforeEach(func() {
				store.DB.Exec(
					"INSERT INTO users (user_id, email) VALUES ($1, $2, $3, $4)",
					userID,
					"lukas.hambsch@gmail.com",
				)

				res, _, _ = Request("DELETE", fmt.Sprintf("%s/%d", userURL, userID), token, nil)
			})

			AfterEach(func() {
				datastore.DeleteUser(userID)
			})

			It("should return status code 200", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})

			It("should delete the user", func() {
				_, err := datastore.GetUser(userID)
				Expect(err).ToNot(BeNil())
			})
		})

		Describe("Unsuccessful DELETE", func() {
			var errRes handlers.APIErrorMessage

			It("should return status code 400 with a message", func() {
				res, data, _ = Request("DELETE", fmt.Sprintf("%s/a", userURL), token, nil)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(errRes.Message).To(Equal(handlers.InvalidUserID))
			})
		})
	})
})
