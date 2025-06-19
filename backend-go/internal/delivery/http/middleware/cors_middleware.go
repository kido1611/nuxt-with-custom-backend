// Package middleware provide basic middleware
// Included CORS, CSRF, Session, and verify Origin
package middleware

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewCorsMiddleware(viper *viper.Viper) fiber.Handler {
	allowedOriginsString := viper.GetString("cors.origins")
	allowedOrigins := strings.Split(allowedOriginsString, ", ")

	allowCredentials := viper.GetBool("cors.allow_credentials")

	allowedMethodsString := viper.GetString("cors.methods")
	allowedHeadersString := viper.GetString("cors.headers")

	return func(ctx *fiber.Ctx) error {
		if isPreflightRequest(ctx) {
			if !originHeader(ctx, allowedOrigins) {
				return fiber.ErrForbidden
			}

			if allowCredentials {
				ctx.Set(fiber.HeaderAccessControlAllowCredentials, "true")
			}

			ctx.Set(fiber.HeaderAccessControlAllowMethods, allowedMethodsString)
			ctx.Set(fiber.HeaderAccessControlAllowHeaders, allowedHeadersString)

			return ctx.SendStatus(http.StatusNoContent)
		}

		err := ctx.Next()

		if originHeader(ctx, allowedOrigins) {
			if allowCredentials {
				ctx.Set(fiber.HeaderAccessControlAllowCredentials, "true")
			}
		}

		return err
	}
}

func isPreflightRequest(ctx *fiber.Ctx) bool {
	method := ctx.Method()
	accessControlRequestMethod := ctx.Get(fiber.HeaderAccessControlRequestMethod)

	fmt.Println(method, accessControlRequestMethod, method == "OPTIONS" && accessControlRequestMethod != "")

	return method == "OPTIONS" && accessControlRequestMethod != ""
}

func originHeader(ctx *fiber.Ctx, allowedOrigins []string) bool {
	ctx.Vary(fiber.HeaderOrigin)

	origin := ctx.Get(fiber.HeaderOrigin, "")
	if origin == "" {
		return false
	}

	if !slices.Contains(allowedOrigins, strings.TrimRight(origin, "/")) {
		return false
	}

	ctx.Set(fiber.HeaderAccessControlAllowOrigin, origin)

	return true
}
