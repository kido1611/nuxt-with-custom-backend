package converter

import (
	"kido1611/notes-backend-go/internal/db/sqlc"
	"kido1611/notes-backend-go/internal/model"
)

func SessionToResponse(session *sqlc.Session) *model.SessionResponse {
	userId := ""
	ipAddress := ""
	userAgent := ""

	if session.UserID.Valid {
		userId = session.UserID.String
	}

	if session.IpAddress.Valid {
		ipAddress = session.IpAddress.String
	}

	if session.UserAgent.Valid {
		ipAddress = session.UserAgent.String
	}

	return &model.SessionResponse{
		ID:        session.ID,
		UserID:    userId,
		CsrfToken: session.CsrfToken,
		IpAddress: ipAddress,
		UserAgent: userAgent,
		ExpiredAt: session.ExpiredAt,
	}
}
