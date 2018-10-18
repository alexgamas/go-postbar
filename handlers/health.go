package handlers

import (
	"encoding/json"
	"fmt"
	"go-postbar/database"
	"go-postbar/model"
	"net/http"
)

// HealthHandler ...
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	elapsed := fmt.Sprintf("%s", database.Ping())

	status := model.Health{Status: "UP", DbPing: elapsed}
	json.NewEncoder(w).Encode(status)
}
