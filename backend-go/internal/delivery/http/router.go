package http

import (
	"database/sql"
	"kido1611/notes-backend-go/internal/delivery/http/controller"
	"kido1611/notes-backend-go/internal/repository"
	"kido1611/notes-backend-go/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Router struct {
	App            *fiber.App
	Validate       *validator.Validate
	Log            *logrus.Logger
	HomeController *controller.HomeController
	AuthController *controller.AuthController
}

func NewRouter(
	DB *sql.DB,
	app *fiber.App,
	validate *validator.Validate,
	log *logrus.Logger,
) *Router {
	userRepository := repository.NewUserRepository()

	userUseCase := usecase.NewUserUseCase(DB, validate, log, userRepository)

	homeController := controller.NewHomeController()
	authController := controller.NewAuthController(log, userUseCase)

	return &Router{
		App:            app,
		HomeController: homeController,
		AuthController: authController,
	}
}

func (r *Router) Setup() {
	r.SetupGuestRoutes()
}

func (r *Router) SetupGuestRoutes() {
	r.App.Get("/", r.HomeController.Index)
	r.App.Post("/auth/register", r.AuthController.Register)
}
