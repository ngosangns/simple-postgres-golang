package user

import (
	"context"

	"simple-postgres-golang/postgres"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	PgxPool *pgxpool.Pool
}

func (s *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*postgres.User, error) {
	conn, err := s.PgxPool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	result, err := postgres.New(conn).GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *UserService) CreateUser(ctx context.Context, params CreateUserParams) (*postgres.User, error) {
	conn, err := s.PgxPool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	result, err := postgres.New(conn).CreateUser(ctx, postgres.CreateUserParams{
		ID:    params.ID,
		Email: params.Email,
	})
	if err != nil {
		return nil, err
	}

	return &result, nil
}
