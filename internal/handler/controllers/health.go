package controllers

import (
	"io"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future, we could check if the database is up and running.
	io.WriteString(w, `{"alive": true}`)
}
