package handlers

import (
	"CardHero/db"
	"CardHero/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	uuid "github.com/satori/go.uuid"
	"io"
	"net/http"
)

func GetFolders(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	user, err := db.GetUser(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	path := chi.URLParam(r, "*")

	if path == "" {
		// Requests for root are sent to / to make for an elegant API
		path = models.RootFolderName
	} else {
		// Append the root folder name to the string
		path = fmt.Sprintf("%s/%s", models.RootFolderName, path)
	}

	fs, err := db.GetFolderStructure(path, *user)
	if err != nil {
		// TODO
	}

	fsJSON, err := json.Marshal(fs)
	if err != nil {
		// TODO
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(fsJSON)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GetCardsFromFolder(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	user, err := db.GetUser(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	folderIDStr := chi.URLParam(r, "folderID")
	folderID, err := uuid.FromString(folderIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid folder ID")
		return
	}

	cards, err := db.GetCardsInFolder(folderID, *user)
	if err != nil {
		http.NotFoundHandler()
		return
	}

	fmt.Println(cards)

	cardsJSON, err := json.Marshal(cards)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(cardsJSON)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
