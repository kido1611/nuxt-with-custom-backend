package middleware

import (
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewVerifyOrigin(viper *viper.Viper) fiber.Handler {
	allowedDomains := viper.GetStringSlice("app.allowed_ui")

	return func(ctx *fiber.Ctx) error {
		origin := ctx.Get("Origin", ctx.Get("Referer", ""))

		origin = strings.Replace(origin, "https://", "", -1)
		origin = strings.Replace(origin, "http://", "", -1)
		origin = strings.TrimSuffix(origin, "/")

		if !slices.Contains(allowedDomains, origin) {
			return fiber.ErrForbidden
		}

		return ctx.Next()
	}
}
