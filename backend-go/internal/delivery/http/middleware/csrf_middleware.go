package middleware

import (
	"kido1611/notes-backend-go/internal/model"
	"slices"

	"github.com/gofiber/fiber/v2"
)

func NewCsrfMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		method := string(ctx.Request().Header.Method())

		notSafeMethods := []string{"POST", "PUT", "PATCH", "DELETE"}

		if slices.Contains(notSafeMethods, method) {
			requestCsrfToken := ctx.Get("X-XSRF-TOKEN", "")
			if requestCsrfToken == "" {
				return fiber.NewError(419, "CSRF Token Missmatch")
			}

			sessionResponse, okSession := ctx.Locals("session").(*model.SessionResponse)
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
		sessionResponse, okSession := ctx.Locals("session").(*model.SessionResponse)
		if okSession && sessionResponse != nil {
			ctx.Set("X-XSRF-TOKEN", sessionResponse.CsrfToken)
		}

		return err
	}
}
