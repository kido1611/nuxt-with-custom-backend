package repository

import (
	"context"
	"kido1611/notes-backend-go/internal/db/sqlc"
)

type NoteRepository struct{}

func NewNoteRepository() *NoteRepository {
	return &NoteRepository{}
}

func (e *NoteRepository) ListNotes(ctx context.Context, query *sqlc.Queries, userId string) ([]sqlc.Note, error) {
	return query.ListUserNotes(ctx, userId)
}

func (e *NoteRepository) GetNoteById(ctx context.Context, query *sqlc.Queries, params sqlc.GetUserNoteParams) (sqlc.Note, error) {
	return query.GetUserNote(ctx, params)
}

func (e *NoteRepository) CreateNote(ctx context.Context, query *sqlc.Queries, params sqlc.CreateUserNoteParams) (sqlc.Note, error) {
	return query.CreateUserNote(ctx, params)
}

func (e *NoteRepository) DeleteNote(ctx context.Context, query *sqlc.Queries, params sqlc.DeleteUserNoteParams) error {
	return query.DeleteUserNote(ctx, params)
}
