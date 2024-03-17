package api

import (

	"net/http"

	"github.com/gorilla/mux"

	"testjun/pkg/repository"
)

type api struct {
	r  *mux.Router
	db *repository.Session
}

func New(router *mux.Router, db *repository.Session) *api {
	return &api{r: router, db: db}
}

func (api *api) FillEndpoints() {
	api.r.HandleFunc("/api/auth/{guid}", api.auth)
	api.r.HandleFunc("/api/refresh/{guid}/{token}", api.refrUs)
}

func (api *api) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, api.r)
}