package handlers

import (
	"encoding/json"
	"fmt"
	"go-postbar/database"
	"go-postbar/model"
	"net/http"
	"runtime"
)

// HealthHandler ...
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// fmt.Println(float64(m.Sys))
	// fmt.Println(float64(m.HeapAlloc))

	elapsed := fmt.Sprintf("%s", database.Ping())

	status := model.Health{
		Status: "UP",
		DbPing: elapsed,
		Mem:    float64(m.TotalAlloc)}
	json.NewEncoder(w).Encode(status)
}
