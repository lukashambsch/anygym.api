package main

import "github.com/lukashambsch/gym-all-over/router"

func main() {
	r := router.Load()

	r.Run()
}
