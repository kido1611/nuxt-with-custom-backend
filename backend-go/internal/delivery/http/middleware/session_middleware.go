package middleware

import (
	"kido1611/notes-backend-go/internal/model"
	"kido1611/notes-backend-go/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func NewSession(log *logrus.Logger, sessionUseCase *usecase.SessionUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		sessionIdCookie := ctx.Cookies("app_session")

		if sessionIdCookie != "" {
			sessionResponse, userResponse, err := sessionUseCase.ValidateSession(ctx.UserContext(), sessionIdCookie)
			if err != nil {
				ctx.ClearCookie("app_session")
				log.Warnf("Failed when validate session: %+v", err)
			} else {
				ctx.Locals("session", sessionResponse)
				ctx.Locals("session_user", userResponse)
			}
		}

		err := ctx.Next()
		return err
	}
}

func NewGuestSession(log *logrus.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		sessionResponse, _ := ctx.Locals("session").(*model.SessionResponse)
		userResponse, okUser := ctx.Locals("session_user").(*model.UserResponse)

		// only deny when session has user
		// there is possibility for session without user for auth/registration
		if okUser && userResponse != nil {
			log.Warnf("User is authenticated: %s", sessionResponse.ID)
			return fiber.ErrForbidden
		}

		return ctx.Next()
	}
}

func NewAuthSession(log *logrus.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userResponse, ok := ctx.Locals("session_user").(*model.UserResponse)
		if !ok {
			log.Warn("Failed fetch session from locals")
			return fiber.ErrUnauthorized
		}

		if userResponse == nil {
			log.Warn("User is unauthenticated")
			return fiber.ErrUnauthorized
		}

		err := ctx.Next()
		return err
	}
}
