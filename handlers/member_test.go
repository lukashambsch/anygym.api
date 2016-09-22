package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

var _ = Describe("Member API", func() {
	var (
		server      *httptest.Server
		memberURL   string
		res         *http.Response
		data        []byte
		contentType string       = "application/json"
		client      *http.Client = &http.Client{}
		badPayload  []byte       = []byte(`{"member_id", 1}`)
	)

	BeforeEach(func() {
		server = httptest.NewServer(router.Load())
		memberURL = fmt.Sprintf("%s%s/members", server.URL, router.V1URLBase)
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("GetMembers endpoint", func() {
		var members []models.Member

		Describe("Successful GET w/o query params", func() {
			BeforeEach(func() {
				res, _ = http.Get(memberURL)
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &members)
			})

			It("should return status code 200", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})

			It("should contain the members", func() {
				Expect(len(members)).To(Equal(2))
			})
		})

		Describe("Successful GET w/ query params", func() {
			It("should return a list of matching members - member_id", func() {
				res, _ = http.Get(fmt.Sprintf("%s?member_id=1", memberURL))
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &members)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(len(members)).To(Equal(1))
			})

			It("should return a matching member - user_id", func() {
				res, _ = http.Get(fmt.Sprintf("%s?user_id=2", memberURL))
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &members)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(len(members)).To(Equal(1))
			})

			It("should return no members with a valid field but no matches", func() {
				res, _ = http.Get(fmt.Sprintf("%s?member_id=10", memberURL))
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &members)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(len(members)).To(Equal(0))
			})

			It("should sort members by the correct field ascending", func() {
				res, _ = http.Get(fmt.Sprintf("%s?sort_order=asc&order_by=member_id", memberURL))
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &members)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(members[0].MemberID).To(Equal(int64(1)))
				Expect(members[1].MemberID).To(Equal(int64(2)))
			})

			It("should sort members by the correct field descending", func() {
				res, _ = http.Get(fmt.Sprintf("%s?sort_order=desc&order_by=member_id", memberURL))
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &members)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(members[0].MemberID).To(Equal(int64(2)))
				Expect(members[1].MemberID).To(Equal(int64(1)))
			})
		})

		Describe("Unsuccessful GET w/ query params", func() {
			var errRes handlers.APIErrorMessage

			It("should return an error with an invalid field as query param", func() {
				res, _ = http.Get(fmt.Sprintf("%s?invalid=test", memberURL))
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusNotFound))
				Expect(errRes.Message).To(Equal("Invalid field in query params."))
			})

			It("should return an error with an invalid field in order_by", func() {
				res, _ = http.Get(fmt.Sprintf("%s?order_by=invalid", memberURL))
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusNotFound))
				Expect(errRes.Message).To(Equal("Invalid order_by field."))
			})

			It("should return an error with an invalid value for sort_order", func() {
				res, _ = http.Get(fmt.Sprintf("%s?order_by=member_id&sort_order=random", memberURL))
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusNotFound))
				Expect(errRes.Message).To(Equal("sort_order must be either 'asc', 'desc', or ''"))
			})
		})
	})

	Describe("GetMember endpoint", func() {
		var (
			member   models.Member
			memberID int64 = 1
		)

		Describe("Successful GET", func() {
			BeforeEach(func() {
				res, _ = http.Get(fmt.Sprintf("%s/%d", memberURL, memberID))
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &member)
			})

			It("should return status code 200", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})

			It("should contain the member in the response", func() {
				Expect(member.MemberID).To(Equal(memberID))
			})
		})

		Describe("Unsuccessful GET", func() {
			var errRes handlers.APIErrorMessage

			Context("Invalid member_id", func() {
				It("should return status code 400 with a message", func() {
					res, _ = http.Get(fmt.Sprintf("%s/asdf", memberURL))
					data, _ = ioutil.ReadAll(res.Body)
					json.Unmarshal(data, &errRes)
					Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
					Expect(errRes.Message).To(Equal(handlers.InvalidMemberID))
				})
			})

			Context("Non existent member_id", func() {
				It("should return status code 404 with a message", func() {
					res, _ = http.Get(fmt.Sprintf("%s/10", memberURL))
					data, _ = ioutil.ReadAll(res.Body)
					json.Unmarshal(data, &errRes)
					Expect(res.StatusCode).To(Equal(http.StatusNotFound))
					Expect(errRes.Message).ToNot(BeEmpty())
				})
			})
		})

	})

	Describe("PostMember endpoint", func() {
		var (
			member  models.Member
			user    *models.User
			payload []byte
		)

		Describe("Successful POST", func() {
			BeforeEach(func() {
				user, _ = datastore.CreateUser(models.User{Email: "test@email.com"})
				payload = []byte(fmt.Sprintf(`{"user_id": %d, "first_name": "Testing", "last_name": "Post"}`, user.UserID))
				res, _ = http.Post(memberURL, contentType, bytes.NewBuffer(payload))
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &member)
			})

			AfterEach(func() {
				datastore.DeleteMember(member.MemberID)
				datastore.DeleteUser(user.UserID)
			})

			It("should return status code 201", func() {
				Expect(res.StatusCode).To(Equal(http.StatusCreated))
			})

			It("should contain the member", func() {
				Expect(member.UserID).To(Equal(user.UserID))
			})

			It("should save the member", func() {
				Expect(member.MemberID).ToNot(Equal(int64(0)))
			})
		})

		Describe("Unsuccessful POST", func() {
			var errRes handlers.APIErrorMessage

			Describe("Bad Request", func() {
				It("should return status code 400 with a message", func() {
					res, _ = http.Post(
						memberURL,
						contentType,
						bytes.NewBuffer(badPayload),
					)
					data, _ = ioutil.ReadAll(res.Body)
					json.Unmarshal(data, &errRes)
					Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
					Expect(errRes.Message).ToNot(BeEmpty())
				})
			})

			Describe("Internal Server Error", func() {
				It("should return status code 500 with a message", func() {
					payload = []byte(`{"user_id": 1}`)
					res, _ = http.Post(memberURL, contentType, bytes.NewBuffer(payload))
					data, _ = ioutil.ReadAll(res.Body)
					json.Unmarshal(data, &errRes)
					Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
					Expect(errRes.Message).ToNot(BeEmpty())
				})
			})
		})
	})

	Describe("PutMember endpoint", func() {
		var (
			member   models.Member
			payload  []byte = []byte(`{"user_id": 2, "first_name": "Kenzie", "last_name": "Hambsch"}`)
			memberID int64  = 2
		)

		Describe("Successful PUT", func() {
			BeforeEach(func() {
				req, _ := http.NewRequest(
					"PUT",
					fmt.Sprintf("%s/%d", memberURL, memberID),
					bytes.NewBuffer(payload),
				)
				req.Header.Set("Content-Type", contentType)

				res, _ = client.Do(req)
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &member)
			})

			AfterEach(func() {
				datastore.UpdateMember(memberID, models.Member{UserID: int64(1), FirstName: "McKenzie", LastName: "Hambsch"})
			})

			It("should return status code 200", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})

			//It("should contain the member", func() {
			//Expect(member.UserID).To(Equal(int64(2)))
			//})

			It("should save the updated member", func() {
				updated, _ := datastore.GetMember(memberID)
				Expect(updated.FirstName).To(Equal("Kenzie"))
			})
		})

		Describe("Unsuccessful PUT", func() {
			var errRes handlers.APIErrorMessage

			It("should return status code 400 with a message", func() {
				req, _ := http.NewRequest(
					"PUT",
					fmt.Sprintf("%s/%d", memberURL, memberID),
					bytes.NewBuffer(badPayload),
				)
				req.Header.Set("Content-Type", contentType)

				res, _ = client.Do(req)
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(errRes.Message).ToNot(BeEmpty())
			})

			It("should return status code 400 with a message", func() {
				req, _ := http.NewRequest(
					"PUT",
					fmt.Sprintf("%s/a", memberURL),
					bytes.NewBuffer(payload),
				)
				req.Header.Set("Content-Type", contentType)

				res, _ = client.Do(req)
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(errRes.Message).To(Equal(handlers.InvalidMemberID))
			})

			It("should return status code 500 with a message", func() {
				req, _ := http.NewRequest(
					"PUT",
					fmt.Sprintf("%s/5000", memberURL),
					bytes.NewBuffer(payload),
				)
				req.Header.Set("Content-Type", contentType)

				res, _ = client.Do(req)
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
				Expect(errRes.Message).ToNot(BeEmpty())
			})
		})
	})

	Describe("DeleteMember endpoint", func() {
		var memberID int64 = 111
		var user *models.User

		Describe("Successful DELETE", func() {
			BeforeEach(func() {
				user, _ = datastore.CreateUser(models.User{Email: "testing@gmail.com"})
				store.DB.Exec(
					"INSERT INTO members (member_id, user_id, first_name, last_name) VALUES ($1, $2, $3, $4)",
					memberID,
					user.UserID,
					"Lukas",
					"Hambsch",
				)

				req, _ := http.NewRequest(
					"DELETE",
					fmt.Sprintf("%s/%d", memberURL, memberID),
					bytes.NewBuffer([]byte(``)),
				)
				req.Header.Set("Content-Type", contentType)

				res, _ = client.Do(req)
			})

			AfterEach(func() {
				datastore.DeleteUser(user.UserID)
			})

			It("should return status code 200", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})

			It("should delete the member", func() {
				_, err := datastore.GetMember(memberID)
				Expect(err).ToNot(BeNil())
			})
		})

		Describe("Unsuccessful DELETE", func() {
			var errRes handlers.APIErrorMessage

			It("should return status code 400 with a message", func() {
				req, _ := http.NewRequest(
					"DELETE",
					fmt.Sprintf("%s/a", memberURL),
					bytes.NewBuffer([]byte(``)),
				)
				req.Header.Set("Content-Type", contentType)

				res, _ = client.Do(req)
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(errRes.Message).To(Equal(handlers.InvalidMemberID))
			})
		})
	})
})
