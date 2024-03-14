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
		Create(context.Context, dto.CreateEmployee) (*entity.Employee, error)
		FindOne(context.Context, dto.FindOneEmployee) (*entity.Employee, error)
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
