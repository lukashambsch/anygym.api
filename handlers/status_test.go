package handlers_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	//"github.com/lukashambsch/gym-all-over/handlers"
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/router"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Status API", func() {
	var (
		server    *httptest.Server
		statusUrl string
		res       *http.Response
		data      []byte
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

			It("should return statuses code 200", func() {
				Expect(res.StatusCode).To(Equal(200))
			})

			It("should contain data", func() {
				Expect(len(statuses)).To(Equal(4))
			})
		})

		//Context("InternalServer Error", func() {
		//var body handlers.APIErrorMessage

		//BeforeEach(func() {
		//res, _ = http.Get(statusUrl)
		//data, _ = ioutil.ReadAll(res.Body)
		//json.Unmarshal(data, &body)
		//})

		//It("should return status code 500", func() {
		//Expect(res.StatusCode).To(Equal(500))
		//})

		//It("should contain message", func() {
		//Expect(body.Message).To(Equal("Error getting status list."))
		//})
		//})
	})

	AfterEach(func() {
		server.Close()
	})
})
