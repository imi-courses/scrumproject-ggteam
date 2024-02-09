package usecase

import (
	"context"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type (
	// User
	User interface {
		FindAll(context.Context) ([]entity.User, error)
	}
	UserRepo interface {
		FindAll(context.Context) ([]entity.User, error)
		Create(context.Context, dto.CreateUser) (*entity.User, error)
		Delete(context.Context, string) error
	}
)
