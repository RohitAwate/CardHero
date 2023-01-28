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
	card1 := models.NewCard(*user, "https://www.youtube.com/watch?v=ErfnhcEV1O8\n\nexplanation of entropy, softmax, cross-entropy\nweather example\nshannon encoding wala manus")
	card2 := models.NewCard(*user, "Beautiful review of banshees of inisherin\nhttps://dmtalkies.com/the-banshees-of-inisherin-ending-explained-2022-film-martin-mcdonagh/")
	card3 := models.NewCard(*user, "Beautiful review of banshees of inisherin\nhttps://dmtalkies.com/the-banshees-of-inisherin-ending-explained-2022-film-martin-mcdonagh/")
	clog.Append(&card1, &card2, &card3)

	json, err := clog.JSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
