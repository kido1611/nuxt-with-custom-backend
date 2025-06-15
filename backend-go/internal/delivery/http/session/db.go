package session

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"kido1611/notes-backend-go/internal/db/helper"
	"kido1611/notes-backend-go/internal/db/sqlc"
	"kido1611/notes-backend-go/internal/model"
	"kido1611/notes-backend-go/internal/model/converter"
	"kido1611/notes-backend-go/internal/repository"
	"time"

	"github.com/sirupsen/logrus"
)

type DbSessionManager struct {
	DB                *sql.DB
	Log               *logrus.Logger
	SessionRepository repository.SessionRepository
}

func NewDbSessionManager(db *sql.DB, log *logrus.Logger, sessionRepository *repository.SessionRepository) *DbSessionManager {
	return &DbSessionManager{
		DB:                db,
		Log:               log,
		SessionRepository: *sessionRepository,
	}
}

func generateRandomToken(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func (s *DbSessionManager) InsertSession(ctx context.Context, user *model.UserResponse) (*model.SessionResponse, error) {
	userId := sql.NullString{
		Valid: false,
	}

	if user != nil {
		userId = sql.NullString{
			Valid:  true,
			String: user.Id,
		}
	}

	data := sqlc.CreateSessionParams{
		ID:        generateRandomToken(32),
		UserID:    userId,
		CsrfToken: generateRandomToken(32),
		IpAddress: sql.NullString{},
		UserAgent: sql.NullString{},
		ExpiredAt: time.Now().Add(24 * time.Hour),
		LastActivityAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	result, err := helper.DbTransaction(s.DB, s.Log, func(query *sqlc.Queries) (*sqlc.Session, error) {
		res, err := s.SessionRepository.CreateSession(ctx, query, data)
		return &res, err
	})
	if err != nil {
		s.Log.Warnf("Failed write session to database: %+v", err)
		return nil, err
	}

	sessionResponse := converter.SessionToResponse(result)

	return sessionResponse, err
}

func (s *DbSessionManager) GetSessionById(ctx context.Context, id string) (*model.SessionResponse, error) {
	sessionData, err := helper.DbTransaction(s.DB, s.Log, func(qtx *sqlc.Queries) (*sqlc.Session, error) {
		response, err := qtx.GetSessionById(ctx, id)

		return &response, err
	})
	if err != nil {
		s.Log.Warnf("Failed get session from database: %+v", err)
		return nil, err
	}

	sessionResponse := converter.SessionToResponse(sessionData)

	return sessionResponse, nil
}

func (s *DbSessionManager) DeleteSession(ctx context.Context, id string) error {
	_, err := helper.DbTransaction(s.DB, s.Log, func(qtx *sqlc.Queries) (*any, error) {
		err := qtx.DeleteSession(ctx, id)
		return nil, err
	})

	return err
}
