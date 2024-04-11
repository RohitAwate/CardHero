package handlers

import (
	"CardHero/db"
	"CardHero/models"
	"CardHero/monitoring"
	"fmt"
	"net/http"
)

func SignInUser(w http.ResponseWriter, r *http.Request) {
	var monitor monitoring.Monitor = monitoring.NewPrintMonitor("api/handlers/user.go#SignInUser()")
	username := r.FormValue("username")
	password := r.FormValue("password")
	user, err := db.GetUserByLoginCredentials(username, password)
	if err != nil {
		errString := fmt.Sprintf("User not found: %s", username)
		LogAndRespond(errString, monitor, monitoring.LogLevelInfo, w, http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, user.Username)

	http.SetCookie(w, &http.Cookie{Name: username})
}

func SignUpUser(w http.ResponseWriter, r *http.Request) {
	var monitor monitoring.Monitor = monitoring.NewPrintMonitor("api/handlers/user.go#SignUpUser()")

	// Username exists?
	username := r.FormValue("username")
	if _, err := db.GetUserByUsername(username); err == nil {
		errString := fmt.Sprintf("Username already exists: %s", username)
		LogAndRespond(errString, monitor, monitoring.LogLevelInfo, w, http.StatusBadRequest)
		return
	}

	// Email exists?
	email := r.FormValue("email")
	if _, err := db.GetUserByEmail(email); err == nil {
		errString := fmt.Sprintf("Account already exists with email: %s", email)
		LogAndRespond(errString, monitor, monitoring.LogLevelInfo, w, http.StatusBadRequest)
		return
	}

	// Insert into users table
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	password := r.FormValue("password")
	user, err := models.NewUser(username, firstName, lastName, email, password)
	if err != nil {
		errString := fmt.Sprintf("Malformed email address: %s", email)
		LogAndRespond(errString, monitor, monitoring.LogLevelInfo, w, http.StatusBadRequest)
		return
	}

	if err = db.SaveUser(*user); err != nil {
		errString := fmt.Sprintf("Error while saving user: %s", err)
		LogAndRespond(errString, monitor, monitoring.LogLevelAlert, w, http.StatusInternalServerError)
		return
	}

	// Create a new root folder for the new user
	root := models.NewRoot(*user)
	if err := db.SaveFolder(&root); err != nil {
		errString := fmt.Sprintf("Error while creating root folder for new user: %s", err)
		monitor.LogError(errString)
	}

	w.WriteHeader(http.StatusCreated)
}
