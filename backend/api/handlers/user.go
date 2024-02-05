package handlers

import (
	"CardHero/db"
	"CardHero/models"
	"CardHero/monitoring"
	"fmt"
	"net/http"
)

func SignUpUser(w http.ResponseWriter, r *http.Request) {
	// Username exists?
	username := r.FormValue("username")
	_, err := db.GetUserByUsername(username)
	var monitor monitoring.Monitor = monitoring.NewPrintMonitor("handlers/user.go#SignUpUser()")
	if err == nil {
		errString := fmt.Sprintf("Username already exists: %s", username)
		LogAndRespond(errString, w, http.StatusBadRequest, monitor, monitoring.LogLevelInfo)
		return
	}

	// Email exists?
	email := r.FormValue("email")
	_, err = db.GetUserByEmail(email)
	if err == nil {
		errString := fmt.Sprintf("Account already exists with email: %s", email)
		LogAndRespond(errString, w, http.StatusBadRequest, monitor, monitoring.LogLevelInfo)
		return
	}

	// Insert into users table
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	password := r.FormValue("password")
	user, err := models.NewUser(username, firstName, lastName, email, password)
	if err != nil {
		errString := fmt.Sprintf("Malformed email address: %s", email)
		LogAndRespond(errString, w, http.StatusBadRequest, monitor, monitoring.LogLevelInfo)
		return
	}

	db.SaveUser(*user)

	// Create a new root folder for the new user
	models.NewRoot(*user)
}
