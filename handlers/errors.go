package handlers

import (
	"net/http"
)

// NotFoundHandler ...
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
