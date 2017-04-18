package main

import (
	"net/http"

	"github.com/lukashambsch/anygym.api/router"
)

func main() {
	r := router.Load()

	http.ListenAndServe(":8080", r)
}
