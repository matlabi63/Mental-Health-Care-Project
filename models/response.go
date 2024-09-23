package models

type Response[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}
