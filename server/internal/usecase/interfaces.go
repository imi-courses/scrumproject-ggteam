package usecase

import (
	"context"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type (
	// Admin
	Admin interface {
		SignUp(context.Context, dto.SignUpAdmin) (*entity.Admin, error)
		FindOne(context.Context, dto.FindOneAdmin) (*entity.Admin, error)
		Delete(context.Context, string) error
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.Admin, error)
	}
	AdminRepo interface {
		Create(context.Context, dto.CreateAdmin) (*entity.Admin, error)
		FindOne(context.Context, dto.FindOneAdmin) (*entity.Admin, error)
		Delete(context.Context, string) error
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.Admin, error)
	}

	// Hash
	Hash interface {
		HashPassword(string) (string, error)
		CheckPasswordHash(string, string) bool
	}

	// Jwt
	Jwt interface {
		CreateToken(dto.TokenPayload, bool) (string, error)
		IsAuthorized(string) (bool, error)
		ExtractIDFromToken(string) (string, error)
	}
)
