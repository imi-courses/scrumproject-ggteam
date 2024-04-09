package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type UseCases struct {
	AdminUseCase    Admin
	EmployeeUseCase Employee
	HashUseCase     Hash
	JwtUseCase      Jwt
}

type (
	Admin interface {
		FindOne(context.Context, entity.Admin) (*entity.Admin, error)
		Delete(context.Context, string) error
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.Admin, error)
	}
	AdminRepo interface {
		Create(context.Context, dto.CreateAdmin) (*entity.Admin, error)
		FindOneById(context.Context, uuid.UUID) (*entity.Admin, error)
		FindOneByEmail(context.Context, string) (*entity.Admin, error)
		FindOneByRefreshToken(context.Context, string) (*entity.Admin, error)
		Delete(context.Context, string) error
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.Admin, error)
	}
	Employee interface {
		Create(context.Context, dto.CreateEmployee) (*entity.Employee, error)
		FindOne(context.Context, entity.Employee) (*entity.Employee, error)
		FindAll(context.Context, dto.Page) ([]entity.Employee, error)
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.Employee, error)
	}
	EmployeeRepo interface {
		Create(context.Context, dto.CreateEmployee) (*entity.Employee, error)
		FindOneById(context.Context, uuid.UUID) (*entity.Employee, error)
		FindOneByEmail(context.Context, string) (*entity.Employee, error)
		FindOneByRefreshToken(context.Context, string) (*entity.Employee, error)
		FindAll(context.Context, dto.Page) ([]entity.Employee, error)
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.Employee, error)
	}
	Hash interface {
		HashPassword(string) (string, error)
		CheckPasswordHash(string, string) bool
	}
	Jwt interface {
		CreateAccessToken(dto.AccessTokenPayload) (string, error)
		CreateRefreshToken(dto.RefreshTokenPayload) (string, error)
		IsTokenValid(string, bool) (bool, error)
		ExtractFromToken(string, string, bool) (string, error)
		CreateTokens(dto.AccessTokenPayload, dto.RefreshTokenPayload) (*dto.Tokens, error)
	}
)
