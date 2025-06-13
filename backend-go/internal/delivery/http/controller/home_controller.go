package controller

import (
	"kido1611/notes-backend-go/internal/model"

	"github.com/gofiber/fiber/v2"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (c *HomeController) Index(ctx *fiber.Ctx) error {
	return ctx.JSON(model.WebResponse[any]{
		Data:    nil,
		Message: "Alive",
	})
}
