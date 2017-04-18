package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/lukashambsch/anygym.api/handlers"
	"github.com/lukashambsch/anygym.api/models"
	"github.com/lukashambsch/anygym.api/router"
	"github.com/lukashambsch/anygym.api/store"
	"github.com/lukashambsch/anygym.api/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Visit API", func() {
	var (
		server     *httptest.Server
		visitURL   string
		res        *http.Response
		data       []byte
		badPayload []byte = []byte(`{"member_id", 1}`)
		token      string
	)

	BeforeEach(func() {
		server = httptest.NewServer(router.Load())
		token, _ = RequestToken(server.URL)
		visitURL = fmt.Sprintf("%s%s/visits", server.URL, router.V1URLBase)
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("GetVisits endpoint", func() {
		var visits []models.Visit

		Describe("Successful GET w/o query params", func() {
			BeforeEach(func() {
				res, data, _ = Request("GET", visitURL, token, nil)
				json.Unmarshal(data, &visits)
			})

			It("should return visit code 200", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})

			It("should contain the visits", func() {
				Expect(len(visits)).To(Equal(5))
			})
		})

		Describe("Successful GET w/ query params", func() {
			It("should return a list of matching visits - member_id", func() {
				res, data, _ = Request("GET", fmt.Sprintf("%s?member_id=1", visitURL), token, nil)
				json.Unmarshal(data, &visits)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(len(visits)).To(Equal(3))
			})

			It("should return a matching visit - visit_id", func() {
				res, data, _ = Request("GET", fmt.Sprintf("%s?visit_id=1", visitURL), token, nil)
				json.Unmarshal(data, &visits)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(len(visits)).To(Equal(1))
			})

			It("should return no visits with a valid field but no matches", func() {
				res, data, _ = Request("GET", fmt.Sprintf("%s?member_id=10", visitURL), token, nil)
				json.Unmarshal(data, &visits)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(len(visits)).To(Equal(0))
			})

			It("should sort visits by the correct field ascending", func() {
				res, data, _ = Request("GET", fmt.Sprintf("%s?sort_order=asc&order_by=member_id", visitURL), token, nil)
				json.Unmarshal(data, &visits)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(visits[0].MemberID).To(Equal(int64(1)))
				Expect(visits[1].MemberID).To(Equal(int64(1)))
				Expect(visits[2].MemberID).To(Equal(int64(1)))
				Expect(visits[3].MemberID).To(Equal(int64(2)))
				Expect(visits[4].MemberID).To(Equal(int64(2)))
			})

			It("should sort visits by the correct field descending", func() {
				res, data, _ = Request("GET", fmt.Sprintf("%s?sort_order=desc&order_by=visit_id", visitURL), token, nil)
				json.Unmarshal(data, &visits)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(visits[0].VisitID).To(Equal(int64(5)))
				Expect(visits[1].VisitID).To(Equal(int64(4)))
				Expect(visits[2].VisitID).To(Equal(int64(3)))
				Expect(visits[3].VisitID).To(Equal(int64(2)))
				Expect(visits[4].VisitID).To(Equal(int64(1)))
			})
		})

		Describe("Unsuccessful GET w/ query params", func() {
			var errRes handlers.APIErrorMessage

			It("should return an error with an invalid field as query param", func() {
				res, data, _ = Request("GET", fmt.Sprintf("%s?invalid=test", visitURL), token, nil)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusNotFound))
				Expect(errRes.Message).To(Equal("Invalid field in query params."))
			})

			It("should return an error with an invalid field in order_by", func() {
				res, data, _ = Request("GET", fmt.Sprintf("%s?order_by=invalid", visitURL), token, nil)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusNotFound))
				Expect(errRes.Message).To(Equal("Invalid order_by field."))
			})

			It("should return an error with an invalid value for sort_order", func() {
				res, data, _ = Request("GET", fmt.Sprintf("%s?order_by=member_id&sort_order=random", visitURL), token, nil)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusNotFound))
				Expect(errRes.Message).To(Equal("sort_order must be either 'asc', 'desc', or ''"))
			})
		})
	})

	Describe("GetVisit endpoint", func() {
		var (
			visit   models.Visit
			visitID int64 = 1
		)

		Describe("Successful GET", func() {
			BeforeEach(func() {
				res, data, _ = Request("GET", fmt.Sprintf("%s/%d", visitURL, visitID), token, nil)
				json.Unmarshal(data, &visit)
			})

			It("should return visit code 200", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})

			It("should contain the visit in the response", func() {
				Expect(visit.VisitID).To(Equal(visitID))
			})
		})

		Describe("Unsuccessful GET", func() {
			var errRes handlers.APIErrorMessage

			Context("Invalid visit_id", func() {
				It("should return visit code 400 with a message", func() {
					res, data, _ = Request("GET", fmt.Sprintf("%s/asdf", visitURL), token, nil)
					json.Unmarshal(data, &errRes)
					Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
					Expect(errRes.Message).To(Equal(handlers.InvalidVisitID))
				})
			})

			Context("Non existent visit_id", func() {
				It("should return visit code 404 with a message", func() {
					res, data, _ = Request("GET", fmt.Sprintf("%s/10", visitURL), token, nil)
					json.Unmarshal(data, &errRes)
					Expect(res.StatusCode).To(Equal(http.StatusNotFound))
					Expect(errRes.Message).ToNot(BeEmpty())
				})
			})
		})

	})

	Describe("PostVisit endpoint", func() {
		var (
			visit   models.Visit
			payload []byte = []byte(`{"member_id": 1, "gym_location_id": 1, "status_id": 1}`)
		)

		Describe("Successful POST", func() {
			BeforeEach(func() {
				res, data, _ = Request("POST", visitURL, token, payload)
				json.Unmarshal(data, &visit)
			})

			AfterEach(func() {
				datastore.DeleteVisit(visit.VisitID)
			})

			It("should return visit code 201", func() {
				Expect(res.StatusCode).To(Equal(http.StatusCreated))
			})

			It("should contain the visit", func() {
				Expect(visit.MemberID).To(Equal(int64(1)))
			})

			It("should save the visit", func() {
				Expect(visit.VisitID).ToNot(Equal(int64(0)))
			})
		})

		Describe("Unsuccessful POST", func() {
			var errRes handlers.APIErrorMessage

			Describe("Bad Request", func() {
				It("should return visit code 400 with a message", func() {
					res, data, _ = Request("POST", visitURL, token, badPayload)
					json.Unmarshal(data, &errRes)
					Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
					Expect(errRes.Message).ToNot(BeEmpty())
				})
			})

			Describe("Internal Server Error", func() {
				It("should return visit code 500 with a message", func() {
					payload = []byte(`{"member_id": 1}`)
					res, data, _ = Request("POST", visitURL, token, payload)
					json.Unmarshal(data, &errRes)
					Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
					Expect(errRes.Message).ToNot(BeEmpty())
				})
			})
		})
	})

	Describe("PutVisit endpoint", func() {
		var (
			visit   models.Visit
			payload []byte = []byte(`{"member_id": 2, "gym_location_id": 1, "status_id": 1}`)
			visitID int64  = 1
		)

		Describe("Successful PUT", func() {
			BeforeEach(func() {
				res, data, _ = Request("PUT", fmt.Sprintf("%s/%d", visitURL, visitID), token, payload)
				json.Unmarshal(data, &visit)
			})

			AfterEach(func() {
				datastore.UpdateVisit(visitID, models.Visit{MemberID: 1})
			})

			It("should return visit code 200", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})

			//It("should contain the visit", func() {
			//Expect(visit.MemberID).To(Equal(int64(2)))
			//})

			It("should save the updated visit", func() {
				updated, _ := datastore.GetVisit(visitID)
				Expect(updated.VisitID).To(Equal(visitID))
			})
		})

		Describe("Unsuccessful PUT", func() {
			var errRes handlers.APIErrorMessage

			It("should return visit code 400 with a message", func() {
				res, data, _ = Request("PUT", fmt.Sprintf("%s/%d", visitURL, visitID), token, badPayload)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(errRes.Message).ToNot(BeEmpty())
			})

			It("should return visit code 400 with a message", func() {
				res, data, _ = Request("PUT", fmt.Sprintf("%s/a", visitURL), token, payload)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(errRes.Message).To(Equal(handlers.InvalidVisitID))
			})

			It("should return visit code 500 with a message", func() {
				res, data, _ = Request("PUT", fmt.Sprintf("%s/5000", visitURL), token, payload)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
				Expect(errRes.Message).ToNot(BeEmpty())
			})
		})
	})

	Describe("DeleteVisit endpoint", func() {
		var visitID int64 = 1

		Describe("Successful DELETE", func() {
			BeforeEach(func() {
				res, _, _ = Request("DELETE", fmt.Sprintf("%s/%d", visitURL, visitID), token, nil)
			})

			AfterEach(func() {
				store.DB.Exec(
					"INSERT INTO visits (visit_id, member_id, gym_location_id, status_id) VALUES ($1, $2, $3, $4)",
					visitID,
					1,
					1,
					1,
				)
			})

			It("should return visit code 200", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})

			It("should delete the visit", func() {
				_, err := datastore.GetVisit(visitID)
				Expect(err).ToNot(BeNil())
			})
		})

		Describe("Unsuccessful DELETE", func() {
			var errRes handlers.APIErrorMessage

			It("should return visit code 400 with a message", func() {
				res, data, _ = Request("DELETE", fmt.Sprintf("%s/a", visitURL), token, nil)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(errRes.Message).To(Equal(handlers.InvalidVisitID))
			})
		})
	})
})
