package handlers

import (
	"CardHero/monitoring"
	"net/http"
)

func LogAndRespond(errString string, monitor monitoring.Monitor, logLevel uint, w http.ResponseWriter, status int) {
	// Respond to request
	w.WriteHeader(status)

	// Log error
	monitoring.LogByLevel(monitor, logLevel, errString)
}
