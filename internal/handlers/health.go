package handlers

import (
	"encoding/json"
	"net/http"
)

type HealthData struct {
	Status string `json:"status"`
}

// handler /Get health
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := HealthData{Status: "ok"}

	response := ApiResponse[HealthData]{
		Success: true,
		Data:    &data,
		Error:   nil,
		Meta:    NewResponseMeta(),
	}

	json.NewEncoder(w).Encode(response)
}

