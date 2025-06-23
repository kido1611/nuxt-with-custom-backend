package middleware

import (
	"kido1611/notes-backend-go/internal/model"
	"kido1611/notes-backend-go/internal/usecase"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewSession(log *logrus.Logger, viper *viper.Viper, sessionUseCase *usecase.SessionUseCase) fiber.Handler {
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

		// update expired and set cookie session if not exist
		sessionResponse, okSession := ctx.Locals("session").(*model.SessionResponse)
		if !okSession || sessionResponse == nil {
			return err
		}

		// currently after access csrf-cookie
		if sessionIdCookie == "" {
			ctx.Cookie(createCookie(viper, sessionResponse.ID, sessionResponse.ExpiredAt))
			return err
		}

		// handle after login
		if sessionIdCookie != sessionResponse.ID {
			ctx.Cookie(createCookie(viper, sessionResponse.ID, sessionResponse.ExpiredAt))
			return err
		}

		newSession, _ := sessionUseCase.UpdateSessionExpired(ctx.UserContext(), sessionResponse)
		if newSession != nil {
			ctx.Cookie(createCookie(viper, newSession.ID, newSession.ExpiredAt))
		}

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

func createCookie(viper *viper.Viper, value string, expires time.Time) *fiber.Cookie {
	cookie := new(fiber.Cookie)
	cookie.Name = "app_session"
	cookie.Value = value
	cookie.Expires = expires
	cookie.HTTPOnly = true
	cookie.SameSite = fiber.CookieSameSiteLaxMode
	cookie.Path = "/"
	cookie.Secure = viper.GetBool("session.secure")
	cookie.Domain = viper.GetString("session.domain")

	return cookie
}
