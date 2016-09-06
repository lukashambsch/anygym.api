package main

import (
	"net/http"

	"github.com/lukashambsch/gym-all-over/router"
)

func main() {
	r := router.Load()

	http.ListenAndServe(":8080", r)
}
