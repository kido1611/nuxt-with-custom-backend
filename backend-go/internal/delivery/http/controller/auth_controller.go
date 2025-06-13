package controller

import (
	"kido1611/notes-backend-go/internal/model"
	"kido1611/notes-backend-go/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuthController struct {
	Log         *logrus.Logger
	UserUseCase *usecase.UserUseCase
}

func NewAuthController(log *logrus.Logger, userUseCase *usecase.UserUseCase) *AuthController {
	return &AuthController{
		Log:         log,
		UserUseCase: userUseCase,
	}
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	request := new(model.RegisterUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed parse request body: %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UserUseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed register user: %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{
		Data: response,
	})
}
