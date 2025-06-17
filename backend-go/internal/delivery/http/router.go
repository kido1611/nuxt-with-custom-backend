package http

import (
	"database/sql"
	"kido1611/notes-backend-go/internal/delivery/http/controller"
	"kido1611/notes-backend-go/internal/delivery/http/middleware"
	"kido1611/notes-backend-go/internal/delivery/http/session"
	"kido1611/notes-backend-go/internal/repository"
	"kido1611/notes-backend-go/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Router struct {
	App            *fiber.App
	Validate       *validator.Validate
	Log            *logrus.Logger
	viper          *viper.Viper
	SessionManager session.SessionManager
	SessionUsecase *usecase.SessionUseCase
	HomeController *controller.HomeController
	AuthController *controller.AuthController
	UserController *controller.UserController
	noteController *controller.NoteController
}

func NewRouter(
	DB *sql.DB,
	app *fiber.App,
	validate *validator.Validate,
	viper *viper.Viper,
	log *logrus.Logger,
) *Router {
	userRepository := repository.NewUserRepository()
	sessionRepository := repository.NewSessionRepository()
	noteRepository := repository.NewNoteRepository()

	sessionManager := session.NewDbSessionManager(DB, log, viper, sessionRepository)

	userUseCase := usecase.NewUserUseCase(DB, validate, log, userRepository)
	sessionUseCase := usecase.NewSessionUseCase(DB, log, sessionManager, userRepository)
	noteUsecase := usecase.NewNoteUsecase(DB, log, validate, noteRepository)

	homeController := controller.NewHomeController()
	authController := controller.NewAuthController(log, viper, userUseCase, sessionManager)
	userController := controller.NewUserController(log)
	noteController := controller.NewNoteController(log, noteUsecase)

	return &Router{
		App:            app,
		Validate:       validate,
		Log:            log,
		viper:          viper,
		SessionManager: sessionManager,
		SessionUsecase: sessionUseCase,
		HomeController: homeController,
		AuthController: authController,
		UserController: userController,
		noteController: noteController,
	}
}

func (r *Router) Setup() {
	r.App.Use(middleware.NewVerifyOrigin(r.viper))
	r.App.Use(middleware.NewSession(r.Log, r.SessionUsecase))
	r.App.Use(middleware.NewCsrfMiddleware())
	r.App.Get("/health", r.HomeController.Index)
	r.App.Get("/sanctum/csrf-cookie", r.AuthController.CsrfToken)
	r.SetupGuestRoutes()
	r.SetupAuthRoutes()
}

func (r *Router) SetupGuestRoutes() {
	// api := r.App.Group("/api/auth", middleware.NewGuestSession(r.Log))
	r.App.Post("/api/auth/register", middleware.NewGuestSession(r.Log), r.AuthController.Register)
	r.App.Post("/api/auth/login", middleware.NewGuestSession(r.Log), r.AuthController.Login)
}

func (r *Router) SetupAuthRoutes() {
	api := r.App.Group("/api/user", middleware.NewAuthSession(r.Log))
	api.Get("/", r.UserController.GetUser)

	r.App.Delete("/api/auth/logout", middleware.NewAuthSession(r.Log), r.AuthController.Logout)

	apiNotes := r.App.Group("/api/notes", middleware.NewAuthSession(r.Log))
	apiNotes.Get("/", r.noteController.ListNotes)
	apiNotes.Post("/", r.noteController.CreateNote)
	apiNotes.Delete("/:noteId", r.noteController.DeleteNote)
}
