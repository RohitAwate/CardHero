package handlers

import (
	"CardHero/db"
	"CardHero/models"
	"CardHero/monitoring"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	uuid "github.com/satori/go.uuid"
)

func GetCards(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	user, err := db.GetUserByUsername(username)
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

func GetCardByID(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	user, err := db.GetUserByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	cardIDStr := chi.URLParam(r, "cardID")
	cardID, err := uuid.FromString(cardIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid folder ID")
		return
	}

	card, err := db.GetCardByID(*user, cardID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cardJSON, err := json.Marshal(card)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(cardJSON)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func AddCard(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	contents := r.FormValue("contents")
	timestampStr := r.FormValue("timestamp")
	timestamp, _ := time.Parse(time.RFC3339, timestampStr)

	user, err := db.GetUserByUsername(username)
	var monitor monitoring.Monitor = monitoring.NewPrintMonitor("api/handlers/card.go#AddCard()")
	if err != nil {
		monitor.LogCaution(fmt.Sprintf("User not found in database: @%s", username))
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

func GetFolderPathByCardID(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	user, err := db.GetUserByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	cardIDStr := chi.URLParam(r, "cardID")
	cardID, err := uuid.FromString(cardIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid folder ID")
		return
	}

	card, err := db.GetCardByID(*user, cardID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	folderPath, err := db.GetFolderPathByID(card.FolderID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	redirectLocation := "/folders"
	for _, folder := range folderPath {
		redirectLocation += "/" + folder
	}
	redirectLocation += "?card=" + cardIDStr

	w.Header().Add("content-type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(redirectLocation))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
