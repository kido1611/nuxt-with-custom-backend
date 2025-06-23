package model

import "time"

type NoteResponse struct {
	ID          string    `json:"id,omitempty"`
	UserID      string    `json:"user_id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	IsVisible   bool      `json:"is_visible,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitzero"`
}

type NoteRequest struct {
	Title       string `json:"title" validate:"required,max=200"`
	Description string `json:"description" validate:"max=2000"`
}
