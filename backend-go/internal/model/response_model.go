package model

type WebResponse[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}
