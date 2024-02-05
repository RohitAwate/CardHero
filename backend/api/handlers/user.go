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
	var monitor monitoring.Monitor = monitoring.NewPrintMonitor("api/handlers/user.go#SignUpUser()")
	username := r.FormValue("username")
	if _, err := db.GetUserByUsername(username); err == nil {
		errString := fmt.Sprintf("Username already exists: %s", username)
		LogAndRespond(errString, w, http.StatusBadRequest, monitor, monitoring.LogLevelInfo)
		return
	}

	// Email exists?
	email := r.FormValue("email")
	if _, err := db.GetUserByEmail(email); err == nil {
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

	if err = db.SaveUser(*user); err != nil {
		errString := fmt.Sprintf("Error while saving user: %s", err)
		LogAndRespond(errString, w, http.StatusInternalServerError, monitor, monitoring.LogLevelAlert)
		return
	}

	// Create a new root folder for the new user
	models.NewRoot(*user)
}
