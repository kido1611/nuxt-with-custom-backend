package repository

import (
	"context"
	"kido1611/notes-backend-go/internal/db/sqlc"
)

type SessionRepository struct{}

func NewSessionRepository() *SessionRepository {
	return &SessionRepository{}
}

func (r *SessionRepository) GetSessionById(ctx context.Context, query *sqlc.Queries, id string) (sqlc.Session, error) {
	return query.GetSessionById(ctx, id)
}

func (r *SessionRepository) DeleteSession(ctx context.Context, query *sqlc.Queries, id string) error {
	return query.DeleteSession(ctx, id)
}

func (r *SessionRepository) CreateSession(ctx context.Context, query *sqlc.Queries, data sqlc.CreateSessionParams) (sqlc.Session, error) {
	return query.CreateSession(ctx, data)
}
