package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"testjun/pkg/service"

	"github.com/gorilla/mux"
)

func (api *api) auth(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		vars := mux.Vars(r)
		lg, ok := vars["guid"]
		if !ok {
			http.Error(w, "No parameter id", http.StatusInternalServerError)
			return
		}
		_, res, _ := api.db.Find(lg)
		fmt.Println(res, lg)
		if !res {
			token, err := service.CreateToken(lg)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = json.NewEncoder(w).Encode(token)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			_, err = api.db.Create(lg, token.RefreshToken)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		} else {
			token, err := service.CreateToken(lg)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = json.NewEncoder(w).Encode(token)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = api.db.Update(lg, token.RefreshToken)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	} else {
		http.Error(w, "Wrong method", http.StatusInternalServerError)
	}
}

func (api *api) refrUs(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println("1**************************")
		vars := mux.Vars(r)
		lg, ok := vars["guid"]
		if !ok {
			http.Error(w, "No parameter guid", http.StatusInternalServerError)
			return
		}
		tk, ok := vars["token"]
		if !ok {
			http.Error(w, "No parameter token", http.StatusInternalServerError)
			return
		}
		fmt.Println("2************************")
		_, res, _ := api.db.Find(lg)
		if !res {
			http.Error(w, "incorrect user", http.StatusInternalServerError)
			return
		}
		fmt.Println("3************************")
		to, err := base64.StdEncoding.DecodeString(tk)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		guid, err := service.ParseRefreshToken(string(to))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(guid.Subject, lg)
		check := service.CheckTokenHash(lg, guid.Subject)
		if !check {
			http.Error(w, "incorrect token", http.StatusInternalServerError)
			return
		}
		token, err := service.CreateToken(lg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = api.db.Update(lg, token.RefreshToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

}
