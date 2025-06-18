package usecase

import (
	"context"
	"database/sql"
	"errors"
	"kido1611/notes-backend-go/internal/db/helper"
	"kido1611/notes-backend-go/internal/db/sqlc"
	"kido1611/notes-backend-go/internal/delivery/http/session"
	"kido1611/notes-backend-go/internal/model"
	"kido1611/notes-backend-go/internal/model/converter"
	"kido1611/notes-backend-go/internal/repository"
	"time"

	"github.com/sirupsen/logrus"
)

type SessionUseCase struct {
	DB             *sql.DB
	Log            *logrus.Logger
	SessionManager session.SessionManager
	UserRepository *repository.UserRepository
}

func NewSessionUseCase(db *sql.DB, log *logrus.Logger, sessionManager session.SessionManager, userRepository *repository.UserRepository) *SessionUseCase {
	return &SessionUseCase{
		DB:             db,
		Log:            log,
		SessionManager: sessionManager,
		UserRepository: userRepository,
	}
}

func (u *SessionUseCase) ValidateSession(ctx context.Context, sessionId string) (*model.SessionResponse, *model.UserResponse, error) {
	sessionResponse, err := u.SessionManager.GetSessionById(ctx, sessionId)
	if err != nil {
		u.Log.Warnf("Session is missing when checking in middleware: %+v", err)
		return nil, nil, err
	}

	if time.Until(sessionResponse.ExpiredAt) < 0*time.Second {
		u.Log.Warnf("Session is expired")
		err := u.SessionManager.DeleteSession(ctx, sessionResponse.ID)
		if err != nil {
			u.Log.Warnf("Failed deleting session: %+v", err)
			return nil, nil, err
		}

		return nil, nil, errors.New("session is expired")
	}

	user, err := helper.DbTransaction(u.DB, u.Log, func(query *sqlc.Queries) (*sqlc.User, error) {
		user, err := query.GetUserById(ctx, sessionResponse.UserID)

		return &user, err
	})
	if err != nil {
		u.Log.Warnf("User missing from database: %+v", err)
		return sessionResponse, nil, nil
	}

	userResponse := converter.UserToResponse(user)
	return sessionResponse, userResponse, nil
}

func (u *SessionUseCase) UpdateSessionExpired(ctx context.Context, sessionResponse *model.SessionResponse) (*model.SessionResponse, error) {
	return u.SessionManager.UpdateExpired(ctx, sessionResponse)
}
