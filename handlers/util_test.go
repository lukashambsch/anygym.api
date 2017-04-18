package handlers_test

import (
	"net/url"

	"github.com/lukashambsch/anygym.api/handlers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handlers utils", func() {
	var fields map[string]string = map[string]string{
		"user_id": "int",
		"email":   "string",
	}
	var params url.Values = url.Values{
		"user_id": []string{"1"},
		"email":   []string{"test@gmail.com"},
	}

	Describe("BuildWhere function", func() {
		var statement string

		Describe("With valid params and fields", func() {
			BeforeEach(func() {
				statement, _ = handlers.BuildWhere(fields, params)
			})

			It("should start with WHERE", func() {
				Expect(statement[:5]).To(Equal("WHERE"))
			})

			It("should contain AND", func() {
				Expect(statement).To(ContainSubstring("AND"))
			})

			It("should construct correct statement - string", func() {
				Expect(statement).To(ContainSubstring("email LIKE '%test@gmail.com%'"))
			})

			It("should construct correct statement - int", func() {
				Expect(statement).To(ContainSubstring("user_id = '1'"))
			})
		})

		Describe("Empty return", func() {
			It("should return empty string if no params", func() {
				params = url.Values{}
				statement, _ = handlers.BuildWhere(fields, params)
				Expect(statement).To(Equal(""))
			})

			It("should return empty string if no fields", func() {
				fields = map[string]string{}
				statement, _ = handlers.BuildWhere(fields, params)
				Expect(statement).To(Equal(""))
			})

			It("should return empty string and error if invalid field", func() {
				params["invalid"] = []string{""}
				statement, err := handlers.BuildWhere(fields, params)
				Expect(statement).To(Equal(""))
				Expect(err.Error()).To(Equal(handlers.InvalidField))
			})
		})
	})
})
