package handlers

import (
	"CardHero/db"
	"CardHero/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

func GetCards(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	user, err := fetchUser(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cards, err := getCards(*user)
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

	fmt.Println(timestampStr)

	timestamp, _ := time.Parse(time.RFC3339, timestampStr)

	user, err := fetchUser(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	card := models.NewCard(*user, contents, timestamp)
	saveCard(card)

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

func fetchUser(username string) (*models.User, error) {
	conn := db.GetConn()

	var user models.User
	err := conn.Find(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func getCards(user models.User) ([]models.Card, error) {
	conn := db.GetConn()

	var cards []models.Card
	err := conn.Find(&cards, "owner_id = ?", user.ID).Error
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func saveCard(card models.Card) {
	conn := db.GetConn()
	conn.Create(&card)
}
