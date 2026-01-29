package handlers

import (
	"encoding/json"
	"net/http"
)

type MetaData struct {
	Service string `json:"service"`
	Version string `json:"version"`
	Runtime string `json:"runtime"`
}

// handler get /meta
func Meta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := MetaData{
		Service: "fs-lab-core-api-go",
		Version: "v0.1.0",
		Runtime: "go",
	}

	response := ApiResponse[MetaData]{
		Success: true,
		Data:    &data,
		Error:   nil,
		Meta:    NewResponseMeta(),
	}

	json.NewEncoder(w).Encode(response)
}
