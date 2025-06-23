package repository

import (
	"context"
	"kido1611/notes-backend-go/internal/db/sqlc"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CountUserByEmail(ctx context.Context, query *sqlc.Queries, email string) (int64, error) {
	return query.CountUserByEmail(ctx, email)
}

func (r *UserRepository) CreateUser(ctx context.Context, query *sqlc.Queries, data sqlc.CreateUserParams) (sqlc.User, error) {
	return query.CreateUser(ctx, data)
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, query *sqlc.Queries, email string) (sqlc.User, error) {
	return query.GetUserByEmail(ctx, email)
}

func (r *UserRepository) GetUserById(ctx context.Context, query *sqlc.Queries, id string) (sqlc.User, error) {
	return query.GetUserById(ctx, id)
}
