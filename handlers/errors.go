package handlers

import (
	"net/http"
)

// NotFoundHandler ...
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Error response 404")
	w.WriteHeader(http.StatusNotFound)
}
