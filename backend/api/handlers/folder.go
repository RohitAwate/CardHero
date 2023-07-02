package handlers

import (
	"CardHero/db"
	"CardHero/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
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

	pathRoot, err := db.ResolveFolder(path, *user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fh, err := db.GetFolderHierarchy(*pathRoot)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fhJSON, err := json.Marshal(fh)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(fhJSON)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GetCardsFromFolderPath(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	user, err := db.GetUser(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	path := chi.URLParam(r, "*")

	if path == "" {
		// Requests for root are interpreted as those for the Default folder
		path = models.DefaultFolderName
	}

	// Append the root folder name to the string
	path = fmt.Sprintf("%s/%s", models.RootFolderName, path)

	folder, err := db.ResolveFolder(path, *user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cards, err := db.GetCardsInFolder(folder.ID, *user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

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
