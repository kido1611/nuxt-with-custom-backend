package middleware

import (
	"kido1611/notes-backend-go/internal/model"
	"slices"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewCsrfMiddleware(viper *viper.Viper) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		method := string(ctx.Request().Header.Method())

		notSafeMethods := []string{"POST", "PUT", "PATCH", "DELETE"}

		sessionResponse, okSession := ctx.Locals("session").(*model.SessionResponse)

		if slices.Contains(notSafeMethods, method) {
			requestCsrfToken := ctx.Get("X-XSRF-TOKEN", "")
			if requestCsrfToken == "" {
				return fiber.NewError(419, "CSRF Token Missmatch")
			}

			if !okSession {
				return fiber.NewError(419, "CSRF Token Missmatch")
			}

			if sessionResponse == nil {
				return fiber.NewError(419, "CSRF Token Missmatch")
			}

			if requestCsrfToken != sessionResponse.CsrfToken {
				return fiber.NewError(419, "CSRF Token Missmatch")
			}
		}

		err := ctx.Next()

		// add csrf heaader
		sessionNewResponse, okNewSession := ctx.Locals("session").(*model.SessionResponse)
		if !okNewSession {
			if sessionResponse != nil {
				// Delete Cookie
				ctx.Cookie(CreateCookie(viper, "XSRF-TOKEN", "", time.Unix(0, 0), false))
			}

			return err
		}

		if sessionNewResponse != nil {
			ctx.Cookie(CreateCookie(viper, "XSRF-TOKEN", sessionNewResponse.CsrfToken, sessionNewResponse.ExpiredAt, false))
		} else {
			// Delete Cookie
			ctx.Cookie(CreateCookie(viper, "XSRF-TOKEN", "", time.Unix(0, 0), false))
		}

		return err
	}
}
