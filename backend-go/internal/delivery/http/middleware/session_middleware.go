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
		sessionIDCookie := ctx.Cookies("app_session")

		if sessionIDCookie != "" {
			sessionResponse, userResponse, err := sessionUseCase.ValidateSession(ctx.UserContext(), sessionIDCookie)
			if err != nil {
				// Delete Cookie
				ctx.Cookie(CreateCookie(viper, "app_session", "", time.Unix(0, 0), true))
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
			if sessionIDCookie != "" {
				// Delete Cookie
				ctx.Cookie(CreateCookie(viper, "app_session", "", time.Unix(0, 0), true))
			}
			return err
		}

		// currently after access csrf-cookie
		if sessionIDCookie == "" {
			ctx.Cookie(CreateCookie(viper, "app_session", sessionResponse.ID, sessionResponse.ExpiredAt, true))
			return err
		}

		// handle after login
		if sessionIDCookie != sessionResponse.ID {
			ctx.Cookie(CreateCookie(viper, "app_session", sessionResponse.ID, sessionResponse.ExpiredAt, true))
			return err
		}

		newSession, _ := sessionUseCase.UpdateSessionExpired(ctx.UserContext(), sessionResponse)
		if newSession != nil {
			ctx.Cookie(CreateCookie(viper, "app_session", newSession.ID, newSession.ExpiredAt, true))
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

func CreateCookie(viper *viper.Viper, name string, value string, expires time.Time, httpOnly bool) *fiber.Cookie {
	cookie := new(fiber.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = expires
	cookie.HTTPOnly = httpOnly
	cookie.SameSite = fiber.CookieSameSiteLaxMode
	cookie.Path = "/"
	cookie.Secure = viper.GetBool("session.secure")

	sessionDomain := viper.GetString("session.domain")
	if sessionDomain != "" {
		cookie.Domain = sessionDomain
	}

	return cookie
}
