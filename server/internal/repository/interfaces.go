package repository

import (
	"context"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type (
	Admin interface {
		Create(context.Context, dto.CreateAdmin) (*entity.Admin, error)
		FindOne(context.Context, dto.FindOneAdmin) (*entity.Admin, error)
		Delete(context.Context, string) error
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.Admin, error)
	}
	Employee interface {
		Create(context.Context, dto.CreateEmployee) (*entity.Employee, error)
		FindOne(context.Context, dto.FindOneEmployee) (*entity.Employee, error)
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.Employee, error)
	}
)
