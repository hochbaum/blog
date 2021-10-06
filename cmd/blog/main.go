package main

import (
	"github.com/gorilla/mux"
	"hochbaum.dev/blog/blog"
)

func main() {
	router := mux.NewRouter()
	storage := blog.NewDummyStorage()

	if err := blog.New(router, storage).Publish(); err != nil {
		panic(err)
	}
}