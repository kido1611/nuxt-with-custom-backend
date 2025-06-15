package controller

import (
	"kido1611/notes-backend-go/internal/model"
	"kido1611/notes-backend-go/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type NoteController struct {
	log         *logrus.Logger
	noteUsecase *usecase.NoteUsecase
}

func NewNoteController(log *logrus.Logger, noteUsecase *usecase.NoteUsecase) *NoteController {
	return &NoteController{
		log:         log,
		noteUsecase: noteUsecase,
	}
}

func (c *NoteController) ListNotes(ctx *fiber.Ctx) error {
	userResponse, ok := ctx.Locals("session_user").(*model.UserResponse)
	if !ok || userResponse == nil {
		c.log.Warnf("User in unauthenticated")
		return fiber.ErrUnauthorized
	}

	notes, err := c.noteUsecase.ListNotes(ctx.UserContext(), userResponse)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(model.WebResponse[[]model.NoteResponse]{
		Data: notes,
	})
}

func (c *NoteController) CreateNote(ctx *fiber.Ctx) error {
	request := new(model.NoteRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.log.Warnf("Failed parse request body: %+v", err)
		return fiber.ErrBadRequest
	}

	userResponse, ok := ctx.Locals("session_user").(*model.UserResponse)
	if !ok || userResponse == nil {
		c.log.Warnf("User in unauthenticated")
		return fiber.ErrUnauthorized
	}

	note, err := c.noteUsecase.CreateNote(ctx.UserContext(), userResponse, request)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.Status(201).JSON(model.WebResponse[*model.NoteResponse]{
		Data: note,
	})
}

func (c *NoteController) DeleteNote(ctx *fiber.Ctx) error {
	userResponse, ok := ctx.Locals("session_user").(*model.UserResponse)
	if !ok || userResponse == nil {
		c.log.Warnf("User in unauthenticated")
		return fiber.ErrUnauthorized
	}
	noteId := ctx.Params("noteId")

	err := c.noteUsecase.DeleteNote(ctx.UserContext(), userResponse, noteId)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.SendStatus(204)
}
