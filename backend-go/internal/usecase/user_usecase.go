package usecase

import (
	"context"
	"database/sql"
	"kido1611/notes-backend-go/internal/db/helper"
	"kido1611/notes-backend-go/internal/db/sqlc"
	"kido1611/notes-backend-go/internal/model"
	"kido1611/notes-backend-go/internal/model/converter"
	"kido1611/notes-backend-go/internal/repository"

	"github.com/alexedwards/argon2id"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UserUseCase struct {
	DB             *sql.DB
	Validate       *validator.Validate
	Log            *logrus.Logger
	UserRepository *repository.UserRepository
}

func NewUserUseCase(DB *sql.DB, validate *validator.Validate, log *logrus.Logger, userRepository *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		DB:             DB,
		Validate:       validate,
		Log:            log,
		UserRepository: userRepository,
	}
}

func (u *UserUseCase) Check(ctx context.Context, request *model.LoginUserRequest) (*model.UserResponse, error) {
	err := u.Validate.Struct(request)
	if err != nil {
		u.Log.Warnf("Invalid request body: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	user, err := helper.DbTransaction(u.DB, u.Log, func(query *sqlc.Queries) (*sqlc.User, error) {
		user, err := u.UserRepository.GetUserByEmail(ctx, query, request.Email)
		return &user, err
	})
	if err != nil {
		u.Log.Warnf("User not found: %+v", err)
		return nil, fiber.ErrUnauthorized
	}

	match, err := argon2id.ComparePasswordAndHash(request.Password, user.Password)
	if err != nil {
		u.Log.Warnf("Failed compare password hash: %+v", err)
		return nil, fiber.ErrUnauthorized
	}

	if !match {
		u.Log.Warnf("User password did not match: %+v", err)
		return nil, fiber.ErrUnauthorized
	}

	userResponse := converter.UserToResponse(user)

	return userResponse, nil
}

func (u *UserUseCase) Create(ctx context.Context, request *model.RegisterUserRequest) (*model.UserResponse, error) {
	err := u.Validate.Struct(request)
	if err != nil {
		u.Log.Warnf("Invalid request body: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	user, err := helper.DbTransaction(u.DB, u.Log, func(query *sqlc.Queries) (*sqlc.User, error) {
		total, err := u.UserRepository.CountUserByEmail(ctx, query, request.Email)
		if err != nil {
			u.Log.Warnf("Failed count user from database: %+v", err)
			return nil, fiber.ErrInternalServerError
		}

		if total > 0 {
			u.Log.Warn("User already exist")
			return nil, fiber.ErrConflict
		}

		hash, err := argon2id.CreateHash(request.Password, argon2id.DefaultParams)
		if err != nil {
			u.Log.Warnf("Failed create password hash: %+v", err)
			return nil, fiber.ErrInternalServerError
		}

		uuid, err := uuid.NewV7()
		if err != nil {
			u.Log.Warnf("Failed generate UUID: %+v", err)
			return nil, fiber.ErrInternalServerError
		}

		uuidString := uuid.String()

		data := sqlc.CreateUserParams{
			ID:       uuidString,
			Name:     request.Name,
			Email:    request.Email,
			Password: hash,
		}

		user, err := u.UserRepository.CreateUser(ctx, query, data)
		return &user, err
	})
	if err != nil {
		u.Log.Warnf("Failed insert user to database: %+v", err)
		return nil, err
	}

	userResponse := converter.UserToResponse(user)

	return userResponse, nil
}
