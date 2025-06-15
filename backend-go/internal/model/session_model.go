package model

import (
	"time"
)

type SessionResponse struct {
	ID             string
	UserID         string
	CsrfToken      string
	IpAddress      string
	UserAgent      string
	ExpiredAt      time.Time
	LastActivityAt time.Time
}
