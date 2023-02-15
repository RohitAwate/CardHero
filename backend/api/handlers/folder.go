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
	fmt.Println(path)

	if path == "" {
		// Requests for root are sent to / to make for an elegant API
		path = models.RootFolderName
	} else {
		// Append the root folder name to the string
		path = fmt.Sprintf("%s/%s", models.RootFolderName, path)
	}

	fs, err := db.GetFolderContents(path, *user)
	if err != nil {

	}

	fsJSON, err := json.Marshal(fs)
	if err != nil {

	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(fsJSON)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
