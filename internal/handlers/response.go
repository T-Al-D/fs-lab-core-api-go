package handlers

import "time"

type ApiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ResponseMeta struct {
	Service   string `json:"service"`
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
}

type ApiResponse[T any] struct {
	Success bool          `json:"success"`
	Data    *T            `json:"data"`
	Error   *ApiError     `json:"error"`
	Meta    ResponseMeta  `json:"meta"`
}

func NewResponseMeta() ResponseMeta {
	return ResponseMeta{
		Service:   "fs-lab-core-api-go",
		Version:   "v0.1.0",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
}
