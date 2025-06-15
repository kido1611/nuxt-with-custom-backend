package converter

import (
	"kido1611/notes-backend-go/internal/db/sqlc"
	"kido1611/notes-backend-go/internal/model"
	"time"
)

func NoteToResponse(note *sqlc.Note) *model.NoteResponse {
	description := ""
	if note.Description.Valid {
		description = note.Description.String
	}

	createdAt := time.Time{}
	if note.CreatedAt.Valid {
		createdAt = note.CreatedAt.Time
	}

	return &model.NoteResponse{
		ID:          note.ID,
		UserID:      note.UserID,
		Title:       note.Title,
		Description: description,
		IsVisible:   note.VisibleAt.Valid,
		CreatedAt:   createdAt,
	}
}
