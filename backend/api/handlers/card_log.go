package handlers

import (
	"CardHero/ch/models"
	"log"
	"net/http"
)

func GetUserLogs(w http.ResponseWriter, r *http.Request) {
	user, err := models.NewUser("Rohit", "Awate", "awate.r@northeastern.edu", "hello123")
	if err != nil {
		log.Fatalln(err)
	}

	clog := models.NewCardLog(*user)
	card1 := models.NewCard(*user, "test card 1")
	card2 := models.NewCard(*user, "test card 2")
	clog.Append(&card1, &card2)

	json, err := clog.JSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
