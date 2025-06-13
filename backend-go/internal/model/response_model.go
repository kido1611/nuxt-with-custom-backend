package model

type WebResponse[T any] struct {
	Data    T      `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}
