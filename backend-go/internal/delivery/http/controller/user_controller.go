package controller

import (
	"kido1611/notes-backend-go/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log *logrus.Logger
}

func NewUserController(log *logrus.Logger) *UserController {
	return &UserController{
		Log: log,
	}
}

func (r *UserController) GetUser(ctx *fiber.Ctx) error {
	userResponse, ok := ctx.Locals("session_user").(*model.UserResponse)
	if !ok {
		r.Log.Warn("Failed parsing locals")
		return fiber.ErrUnauthorized
	}

	if userResponse == nil {
		r.Log.Warn("UserResponse is nil")
		return fiber.ErrUnauthorized
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{
		Data: userResponse,
	})
}
