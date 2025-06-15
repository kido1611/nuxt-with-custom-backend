package converter

import (
	"kido1611/notes-backend-go/internal/db/sqlc"
	"kido1611/notes-backend-go/internal/model"
)

func UserToResponse(user *sqlc.User) *model.UserResponse {
	return &model.UserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Time.UnixMilli(),
	}
}
