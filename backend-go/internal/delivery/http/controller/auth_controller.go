package controller

import (
	localSession "kido1611/notes-backend-go/internal/delivery/http/session"
	"kido1611/notes-backend-go/internal/model"
	"kido1611/notes-backend-go/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuthController struct {
	Log            *logrus.Logger
	UserUseCase    *usecase.UserUseCase
	SessionManager localSession.SessionManager
}

func NewAuthController(log *logrus.Logger, userUseCase *usecase.UserUseCase, sessionManager localSession.SessionManager) *AuthController {
	return &AuthController{
		Log:            log,
		UserUseCase:    userUseCase,
		SessionManager: sessionManager,
	}
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	request := new(model.LoginUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed parse request body: %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UserUseCase.Check(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed register user: %+v", err)
		return err
	}

	// Delete old session first,
	// use this approach to prevent session fixation
	sessionResponse, okSession := ctx.Locals("session").(*model.SessionResponse)
	if okSession && sessionResponse != nil {
		c.SessionManager.DeleteSession(ctx.UserContext(), sessionResponse.ID)
	}

	// Create new session
	session, err := c.SessionManager.InsertSession(ctx.UserContext(), response)
	if err != nil {
		c.Log.Warnf("Failed creating session: %+v", err)
		return err
	}

	// set locals to add csrf token
	ctx.Locals("session", session)

	return ctx.JSON(model.WebResponse[*model.UserResponse]{
		Data: response,
	})
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

	// delete old session too to prevent session fixation
	sessionResponse, okSession := ctx.Locals("session").(*model.SessionResponse)
	if okSession && sessionResponse != nil {
		c.SessionManager.DeleteSession(ctx.UserContext(), sessionResponse.ID)
		ctx.Locals("session", nil)
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse[*model.UserResponse]{
		Data: response,
	})
}

func (c *AuthController) CsrfToken(ctx *fiber.Ctx) error {
	sessionResponse, okSession := ctx.Locals("session").(*model.SessionResponse)
	if okSession && sessionResponse != nil {
		return ctx.SendStatus(204)
	}

	session, err := c.SessionManager.InsertSession(ctx.UserContext(), nil)
	if err != nil {
		c.Log.Warnf("Failed creating session: %+v", err)
		return err
	}

	// set locals to add csrf token
	ctx.Locals("session", session)

	return ctx.SendStatus(204)
}

func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	sessionResponse, okSession := ctx.Locals("session").(*model.SessionResponse)
	if okSession && sessionResponse == nil {
		return ctx.SendStatus(204)
	}

	err := c.SessionManager.DeleteSession(ctx.UserContext(), sessionResponse.ID)
	if err != nil {
		c.Log.Warnf("Failed deleting session: %+v", err)
		return err
	}
	ctx.Locals("session", nil)

	// ctx.ClearCookie("app_session")
	return ctx.SendStatus(204)
}
