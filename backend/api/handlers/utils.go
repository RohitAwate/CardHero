package handlers

import (
	"CardHero/monitoring"
	"net/http"
)

func LogAndRespond(errString string, w http.ResponseWriter, status int, monitor monitoring.Monitor, logLevel uint) {
	// Respond to request
	w.WriteHeader(status)

	// Log error
	monitoring.LogByLevel(monitor, logLevel, errString)
}
