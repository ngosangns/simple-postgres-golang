package user

import "github.com/google/uuid"

type CreateUserParams struct {
	Email string
	ID    uuid.UUID
}
