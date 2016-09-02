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

var _ = Describe("Status API", func() {
	var (
		server      *httptest.Server
		statusURL   string
		res         *http.Response
		data        []byte
		contentType string       = "application/json"
		client      *http.Client = &http.Client{}
		badPayload  []byte       = []byte(`{"status_name", "Status Name"}`)
	)

	BeforeEach(func() {
		server = httptest.NewServer(router.Load())
		statusURL = fmt.Sprintf("%s%s/statuses", server.URL, router.V1URLBase)
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("GetStatuses endpoint", func() {
		var statuses []models.Status

		Describe("Successful GET w/o query params", func() {
			BeforeEach(func() {
				res, _ = http.Get(statusURL)
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &statuses)
			})

			It("should return status code 200", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})

			It("should contain the statuses", func() {
				Expect(len(statuses)).To(Equal(4))
			})
		})

		Describe("Successful GET w/ query params", func() {
			It("should return a list of matching statuses - status_name", func() {
				res, _ = http.Get(fmt.Sprintf("%s?status_name=Denied", statusURL))
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &statuses)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(len(statuses)).To(Equal(2))
			})

			It("should return a list of matching statuses - status_name - partial match", func() {
				correct := []models.Status{
					models.Status{StatusId: 1, StatusName: "Pending"},
				}
				res, _ = http.Get(fmt.Sprintf("%s?status_name=Pend", statusURL))
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &statuses)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(statuses).To(Equal(correct))
			})

			It("should return a matching status - status_id", func() {
				res, _ = http.Get(fmt.Sprintf("%s/?status_id=1", statusURL))
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &statuses)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(len(statuses)).To(Equal(1))
			})

			It("should return no statuses with a valid field but no matches", func() {
				res, _ = http.Get(fmt.Sprintf("%s?status_name=Testing", statusURL))
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &statuses)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(len(statuses)).To(Equal(0))
			})

			It("should sort statuses by the correct field ascending", func() {
				res, _ = http.Get(fmt.Sprintf("%s?sort_order=asc&order_by=status_name", statusURL))
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &statuses)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(statuses[0].StatusName).To(Equal("Approved"))
				Expect(statuses[1].StatusName).To(Equal("Denied - Banned"))
				Expect(statuses[2].StatusName).To(Equal("Denied - Identity"))
				Expect(statuses[3].StatusName).To(Equal("Pending"))
			})

			It("should sort statuses by the correct fStatusNameStatusNameStatusNameield descending", func() {
				res, _ = http.Get(fmt.Sprintf("%s?sort_order=desc&order_by=status_id", statusURL))
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &statuses)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(statuses[0].StatusId).To(Equal(int64(4)))
				Expect(statuses[1].StatusId).To(Equal(int64(3)))
				Expect(statuses[2].StatusId).To(Equal(int64(2)))
				Expect(statuses[3].StatusId).To(Equal(int64(1)))
			})
		})

		Describe("Unsuccessful GET w/ query params", func() {
			var errRes handlers.APIErrorMessage

			It("should return an error with an invalid field as query param", func() {
				res, _ = http.Get(fmt.Sprintf("%s?invalid=test", statusURL))
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusNotFound))
				Expect(errRes.Message).To(Equal("Invalid field in query params."))
			})

			It("should return an error with an invalid field in order_by", func() {
				res, _ = http.Get(fmt.Sprintf("%s?order_by=invalid", statusURL))
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusNotFound))
				Expect(errRes.Message).To(Equal("Invalid order_by field."))
			})

			It("should return an error with an invalid value for sort_order", func() {
				res, _ = http.Get(fmt.Sprintf("%s?order_by=status_name&sort_order=random", statusURL))
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusNotFound))
				Expect(errRes.Message).To(Equal("sort_order must be either 'asc', 'desc', or ''"))
			})
		})
	})

	Describe("GetStatus endpoint", func() {
		var (
			status   models.Status
			statusId int64 = 1
		)

		Describe("Successful GET", func() {
			BeforeEach(func() {
				res, _ = http.Get(fmt.Sprintf("%s/%d", statusURL, statusId))
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &status)
			})

			It("should return status code 200", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})

			It("should contain the status in the response", func() {
				Expect(status.StatusId).To(Equal(statusId))
			})
		})

		Describe("Unsuccessful GET", func() {
			var errRes handlers.APIErrorMessage

			Context("Invalid status_id", func() {
				It("should return status code 400 with a message", func() {
					res, _ = http.Get(fmt.Sprintf("%s/asdf", statusURL))
					data, _ = ioutil.ReadAll(res.Body)
					json.Unmarshal(data, &errRes)
					Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
					Expect(errRes.Message).To(Equal(handlers.InvalidStatusId))
				})
			})

			Context("Non existent status_id", func() {
				It("should return status code 404 with a message", func() {
					res, _ = http.Get(fmt.Sprintf("%s/10", statusURL))
					data, _ = ioutil.ReadAll(res.Body)
					json.Unmarshal(data, &errRes)
					Expect(res.StatusCode).To(Equal(http.StatusNotFound))
					Expect(errRes.Message).ToNot(BeEmpty())
				})
			})
		})

	})

	Describe("PostStatus endpoint", func() {
		var (
			status  models.Status
			payload []byte = []byte(`{"status_name": "New Status"}`)
		)

		Describe("Successful POST", func() {
			BeforeEach(func() {
				res, _ = http.Post(fmt.Sprintf("%s", statusURL), contentType, bytes.NewBuffer(payload))
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &status)
			})

			AfterEach(func() {
				datastore.DeleteStatus(status.StatusId)
			})

			It("should return status code 201", func() {
				Expect(res.StatusCode).To(Equal(http.StatusCreated))
			})

			It("should contain the status", func() {
				Expect(status.StatusName).To(Equal("New Status"))
			})

			It("should save the status", func() {
				Expect(status.StatusId).ToNot(Equal(0))
			})
		})

		Describe("Unsuccessful POST", func() {
			var errRes handlers.APIErrorMessage

			Describe("Bad Request", func() {
				It("should return status code 400 with a message", func() {
					res, _ = http.Post(
						fmt.Sprintf("%s", statusURL),
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
					payload = []byte(`{"status_name": "Pending"}`)
					res, _ = http.Post(fmt.Sprintf("%s", statusURL), contentType, bytes.NewBuffer(payload))
					data, _ = ioutil.ReadAll(res.Body)
					json.Unmarshal(data, &errRes)
					Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
					Expect(errRes.Message).ToNot(BeEmpty())
				})
			})
		})
	})

	Describe("PutStatus endpoint", func() {
		var (
			status   models.Status
			payload  []byte = []byte(`{"status_name": "Updated"}`)
			statusId int64  = 1
		)

		Describe("Successful PUT", func() {
			BeforeEach(func() {
				req, _ := http.NewRequest(
					"PUT",
					fmt.Sprintf("%s/%d", statusURL, statusId),
					bytes.NewBuffer(payload),
				)
				req.Header.Set("Content-Type", contentType)

				res, _ = client.Do(req)
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &status)
			})

			AfterEach(func() {
				datastore.UpdateStatus(statusId, models.Status{StatusName: "Pending"})
			})

			It("should return status code 200", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})

			It("should contain the status", func() {
				Expect(status.StatusName).To(Equal("Updated"))
			})

			It("should save the updated status", func() {
				updated, _ := datastore.GetStatus(statusId)
				Expect(updated.StatusId).To(Equal(statusId))
			})
		})

		Describe("Unsuccessful PUT", func() {
			var errRes handlers.APIErrorMessage

			It("should return status code 400 with a message", func() {
				req, _ := http.NewRequest(
					"PUT",
					fmt.Sprintf("%s/%d", statusURL, statusId),
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
					fmt.Sprintf("%s/a", statusURL),
					bytes.NewBuffer(payload),
				)
				req.Header.Set("Content-Type", contentType)

				res, _ = client.Do(req)
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(errRes.Message).To(Equal(handlers.InvalidStatusId))
			})

			It("should return status code 500 with a message", func() {
				req, _ := http.NewRequest(
					"PUT",
					fmt.Sprintf("%s/5", statusURL),
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

	Describe("DeleteStatus endpoint", func() {
		var statusId int64 = 1

		Describe("Successful DELETE", func() {
			BeforeEach(func() {
				req, _ := http.NewRequest(
					"DELETE",
					fmt.Sprintf("%s/%d", statusURL, statusId),
					bytes.NewBuffer([]byte(``)),
				)
				req.Header.Set("Content-Type", contentType)

				res, _ = client.Do(req)
			})

			AfterEach(func() {
				store.DB.QueryRow(
					"INSERT INTO statuses (status_id, status_name) VALUES ($1, $2)",
					statusId,
					"Pending",
				)
			})

			It("should return status code 200", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})

			It("should delete the status", func() {
				_, err := datastore.GetStatus(statusId)
				Expect(err).ToNot(BeNil())
			})
		})

		Describe("Unsuccessful DELETE", func() {
			var errRes handlers.APIErrorMessage

			It("should return status code 400 with a message", func() {
				req, _ := http.NewRequest(
					"DELETE",
					fmt.Sprintf("%s/a", statusURL),
					bytes.NewBuffer([]byte(``)),
				)
				req.Header.Set("Content-Type", contentType)

				res, _ = client.Do(req)
				data, _ = ioutil.ReadAll(res.Body)
				json.Unmarshal(data, &errRes)
				Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(errRes.Message).To(Equal(handlers.InvalidStatusId))
			})
		})
	})
})
