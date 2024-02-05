package handlers

import (
	"CardHero/db"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	query := r.URL.Query().Get("query")

	user, err := db.GetUserByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	results, err := db.Search(query, *user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resultsJSON, err := json.Marshal(results)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resultsJSON)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
