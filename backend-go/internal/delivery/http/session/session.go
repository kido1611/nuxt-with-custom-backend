package session

import (
	"context"
	"kido1611/notes-backend-go/internal/model"
)

type SessionManager interface {
	GetSessionById(ctx context.Context, id string) (*model.SessionResponse, error)
	DeleteSession(ctx context.Context, id string) error
	InsertSession(ctx context.Context, user *model.UserResponse) (*model.SessionResponse, error)
}
