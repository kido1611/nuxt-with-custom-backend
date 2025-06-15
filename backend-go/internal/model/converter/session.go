package converter

import (
	"kido1611/notes-backend-go/internal/db/sqlc"
	"kido1611/notes-backend-go/internal/model"
	"time"
)

func SessionToResponse(session *sqlc.Session) *model.SessionResponse {
	userId := ""
	ipAddress := ""
	userAgent := ""
	lastActivityAt := time.Now()

	if session.UserID.Valid {
		userId = session.UserID.String
	}

	if session.IpAddress.Valid {
		ipAddress = session.IpAddress.String
	}

	if session.UserAgent.Valid {
		ipAddress = session.UserAgent.String
	}

	if session.LastActivityAt.Valid {
		lastActivityAt = session.LastActivityAt.Time
	}

	return &model.SessionResponse{
		ID:             session.ID,
		UserID:         userId,
		CsrfToken:      session.CsrfToken,
		IpAddress:      ipAddress,
		UserAgent:      userAgent,
		ExpiredAt:      session.ExpiredAt,
		LastActivityAt: lastActivityAt,
	}
}
