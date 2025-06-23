package usecase

import (
	"context"
	"database/sql"
	"kido1611/notes-backend-go/internal/db/helper"
	"kido1611/notes-backend-go/internal/db/sqlc"
	"kido1611/notes-backend-go/internal/model"
	"kido1611/notes-backend-go/internal/model/converter"
	"kido1611/notes-backend-go/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type NoteUsecase struct {
	DB             *sql.DB
	Log            *logrus.Logger
	validate       *validator.Validate
	NoteRepository *repository.NoteRepository
}

func NewNoteUsecase(db *sql.DB, log *logrus.Logger, validate *validator.Validate, noteRepository *repository.NoteRepository) *NoteUsecase {
	return &NoteUsecase{
		DB:             db,
		Log:            log,
		validate:       validate,
		NoteRepository: noteRepository,
	}
}

func (u *NoteUsecase) ListNotes(ctx context.Context, userResponse *model.UserResponse) ([]model.NoteResponse, error) {
	notes, err := helper.DbTransaction(u.DB, u.Log, func(query *sqlc.Queries) ([]sqlc.Note, error) {
		notes, err := query.ListUserNotes(ctx, userResponse.Id)

		return notes, err
	})
	if err != nil {
		u.Log.Warnf("Failed fetch user notes: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]model.NoteResponse, len(notes))

	for i, note := range notes {
		responses[i] = *converter.NoteToResponse(&note)
	}

	return responses, nil
}

func (u *NoteUsecase) CreateNote(ctx context.Context, userResponse *model.UserResponse, request *model.NoteRequest) (*model.NoteResponse, error) {
	err := u.validate.Struct(request)
	if err != nil {
		u.Log.Warnf("Invalid request body: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	note, err := helper.DbTransaction(u.DB, u.Log, func(query *sqlc.Queries) (*sqlc.Note, error) {
		uuid, err := uuid.NewV7()
		if err != nil {
			u.Log.Warnf("Failed generate UUID: %+v", err)
			return nil, fiber.ErrInternalServerError
		}

		description := sql.NullString{
			Valid: false,
		}

		if len(request.Description) > 0 {
			description = sql.NullString{
				Valid:  true,
				String: request.Description,
			}
		}

		note, err := query.CreateUserNote(ctx, sqlc.CreateUserNoteParams{
			ID:          uuid.String(),
			UserID:      userResponse.Id,
			Title:       request.Title,
			Description: description,
			VisibleAt: sql.NullTime{
				Valid: false,
			},
		})

		return &note, err
	})
	if err != nil {
		u.Log.Warnf("Failed insert new note: %+v", err)
	}

	noteResponse := converter.NoteToResponse(note)

	return noteResponse, nil
}

func (u *NoteUsecase) DeleteNote(ctx context.Context, userResponse *model.UserResponse, noteId string) error {
	_, err := helper.DbTransaction(u.DB, u.Log, func(query *sqlc.Queries) (*sqlc.Note, error) {
		err := query.DeleteUserNote(ctx, sqlc.DeleteUserNoteParams{
			ID:     noteId,
			UserID: userResponse.Id,
		})

		return nil, err
	})
	if err != nil {
		u.Log.Warnf("Failed delete note: %+v", err)
	}

	return err
}
