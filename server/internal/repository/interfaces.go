package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type (
	Admin interface {
		Create(context.Context, dto.CreateAdmin) (*entity.Admin, error)
		FindOneById(context.Context, uuid.UUID) (*entity.Admin, error)
		FindOneByEmail(context.Context, string) (*entity.Admin, error)
		FindOneByRefreshToken(context.Context, string) (*entity.Admin, error)
		Delete(context.Context, string) error
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.Admin, error)
	}
	Employee interface {
		Create(context.Context, dto.CreateEmployee) (*entity.Employee, error)
		FindOneById(context.Context, uuid.UUID) (*entity.Employee, error)
		FindOneByEmail(context.Context, string) (*entity.Employee, error)
		FindOneByRefreshToken(context.Context, string) (*entity.Employee, error)
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.Employee, error)
	}
)
