package usecase

import (
	"context"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type (
	Admin interface {
		FindOne(context.Context, dto.FindOneAdmin) (*entity.Admin, error)
		Delete(context.Context, string) error
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.Admin, error)
	}
	Employee interface {
		SignUp(context.Context, dto.CreateEmployee) (*entity.Employee, error)
		FindOne(context.Context, dto.FindOneEmployee) (*entity.Employee, error)
	}
	Hash interface {
		HashPassword(string) (string, error)
		CheckPasswordHash(string, string) bool
	}
	Jwt interface {
		CreateToken(dto.TokenPayload, bool) (string, error)
		IsAuthorized(string) (bool, error)
		ExtractIDFromToken(string) (string, error)
	}
)
