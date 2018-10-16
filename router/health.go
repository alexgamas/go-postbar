package router

import (
	"encoding/json"
	"go-postbar/model"
	"net/http"

	"github.com/gorilla/mux"
)

// HealthHandle ...
func HealthHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	status := model.Health{Status: "UP"}
	json.NewEncoder(w).Encode(status)
}

// HealthRouter HealthRouter
func HealthRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HealthHandle)
	return r
}
