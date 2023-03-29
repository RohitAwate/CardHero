package handlers

import (
	"CardHero/db"
	"CardHero/models"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

func GetCards(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	user, err := db.GetUser(username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	cards, err := db.GetCardsBy(*user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var status int
	if len(cards) == 0 {
		status = http.StatusNotFound
	} else {
		status = http.StatusOK
	}

	cardsJson, err := json.Marshal(cards)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(cardsJson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func AddCard(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	contents := r.FormValue("contents")
	timestampStr := r.FormValue("timestamp")
	timestamp, _ := time.Parse(time.RFC3339, timestampStr)

	user, err := db.GetUser(username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	card := models.NewCard(*user, contents, timestamp)
	go db.IngestCard(card, *user)

	cardJson, err := json.Marshal(card)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(cardJson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
