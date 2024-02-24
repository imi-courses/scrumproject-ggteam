package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type (
	// Admin
	Admin interface {
		SignUp(context.Context, dto.SignUpAdmin) (*entity.Admin, error)
		SignIn(context.Context, uuid.UUID) (*entity.Admin, error)
		FindOne(context.Context, uuid.UUID) (*entity.Admin, error)
		Delete(context.Context, string) error
	}
	AdminRepo interface {
		Create(context.Context, dto.CreateAdmin) (*entity.Admin, error)
		FindOne(context.Context, uuid.UUID) (*entity.Admin, error)
		Delete(context.Context, string) error
	}

	// Hash
	Hash interface {
		HashPassword(string) (string, error)
		CheckPasswordHash(string, string) bool
	}
)
