package middleware

import (
	"net/url"
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewVerifyOrigin(viper *viper.Viper) fiber.Handler {
	// use cors origins config
	allowedOriginsString := viper.GetString("cors.origins")
	allowedOrigins := strings.Split(allowedOriginsString, ", ")

	return func(ctx *fiber.Ctx) error {
		origin := ctx.Get("Origin", ctx.Get("Referer", ""))
		origin = GetOriginFromURL(origin)

		if !slices.Contains(allowedOrigins, strings.TrimRight(origin, "/")) {
			return fiber.ErrForbidden
		}

		return ctx.Next()
	}
}

func GetOriginFromURL(origin string) string {
	originURL, err := url.Parse(origin)
	if err != nil {
		return ""
	}

	return originURL.Scheme + "://" + originURL.Host
}
