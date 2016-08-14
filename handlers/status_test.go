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
		server    *httptest.Server
		statusUrl string
		res       *http.Response
		data      []byte
        contentType string = "application/json"
        client *http.Client = &http.Client{}
	)

	BeforeEach(func() {
		server = httptest.NewServer(router.Load())
	})

	Describe("GetStatuses endpoint", func() {
		var statuses []models.Status

		BeforeEach(func() {
			statusUrl = fmt.Sprintf("%s/statuses", server.URL)
		})

		Context("Successful GET", func() {
			BeforeEach(func() {
				res, _ = http.Get(statusUrl)
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
	})

	Describe("GetStatus endpoint", func() {
		var (
			status   models.Status
			statusId int64 = 1
		)

        Describe("Successful GET", func() {
            BeforeEach(func() {
                res, _ = http.Get(fmt.Sprintf("%s/statuses/%d", server.URL, statusId))
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

            BeforeEach(func() {
                res, _ = http.Get(fmt.Sprintf("%s/statuses/asdf", server.URL))
                data, _ = ioutil.ReadAll(res.Body)
                json.Unmarshal(data, &errRes)
            })

            It("should return status code 400", func() {
                Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
            })

            It("should contain the message in the response", func() {
                Expect(errRes.Message).To(Equal(handlers.InvalidStatusId))
            })
        })

	})

	Describe("PostStatus endpoint", func() {
		var (
			status  models.Status
			payload []byte = []byte(`{"status_name": "New Status"}`)
		)

		BeforeEach(func() {
			res, _ = http.Post(fmt.Sprintf("%s/statuses", server.URL), contentType, bytes.NewBuffer(payload))
			data, _ = ioutil.ReadAll(res.Body)
			json.Unmarshal(data, &status)
		})

		Context("Successful POST", func() {
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

		AfterEach(func() {
			datastore.DeleteStatus(status.StatusId)
		})
	})

    Describe("PutStatus endpoing", func() {
        var (
            status models.Status
            payload []byte = []byte(`{"status_name": "Updated"}`)
            statusId int64 = 1
        )

        BeforeEach(func() {
            req, _ := http.NewRequest(
                "PUT",
                fmt.Sprintf("%s/statuses/%d", server.URL, statusId),
                bytes.NewBuffer(payload),
            )
            req.Header.Set("Content-Type", contentType)

            res, _ = client.Do(req)
            data, _ = ioutil.ReadAll(res.Body)
            json.Unmarshal(data, &status)
        })

        Context("Successful PUT", func() {
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

        AfterEach(func() {
            datastore.UpdateStatus(statusId, models.Status{StatusName: "Pending"})
        })
    })

    Describe("DeleteStatus endpoint", func() {
        var statusId int64 = 1

        BeforeEach(func() {
            req, _ := http.NewRequest(
                "DELETE",
                fmt.Sprintf("%s/statuses/%d", server.URL, statusId),
                bytes.NewBuffer([]byte(``)),
            )
            req.Header.Set("Content-Type", contentType)

            res, _ = client.Do(req)
        })

        Context("Successful DELETE", func() {
            It("should return status code 200", func() {
                Expect(res.StatusCode).To(Equal(http.StatusOK))
            })

            It("should delete the status", func() {
                _, err := datastore.GetStatus(statusId)
                Expect(err).ToNot(BeNil())
            })
        })

		AfterEach(func() {
			store.DB.QueryRow(
				"INSERT INTO statuses (status_id, status_name) VALUES ($1, $2)",
				statusId,
				"Pending",
			)
		})
    })

	AfterEach(func() {
		server.Close()
	})
})
