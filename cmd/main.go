package main

import (
	"log"

	"testjun/pkg/api"
	"testjun/pkg/repository"

	"github.com/gorilla/mux"
)
func main () {
	db, err := repository.New("mongodb://127.0.0.1:27017")
	if err != nil {
		log.Fatal(err.Error())
	}
	api := api.New(mux.NewRouter(), db)
	api.FillEndpoints()
	log.Fatal(api.ListenAndServe("localhost:8080"))
}