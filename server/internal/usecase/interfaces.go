package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type UseCases struct {
	AdminUseCase      Admin
	EmployeeUseCase   Employee
	ClientUseCase     Client
	RealEstateUseCase RealEstate
	HashUseCase       Hash
	JwtUseCase        Jwt
}

type (
	// Users
	Admin interface {
		FindOne(context.Context, entity.Admin) (*entity.Admin, error)
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.Admin, error)
	}
	AdminRepo interface {
		Create(context.Context, dto.CreateAdmin) (*entity.Admin, error)
		FindOneById(context.Context, uuid.UUID) (*entity.Admin, error)
		FindOneByEmail(context.Context, string) (*entity.Admin, error)
		FindOneByRefreshToken(context.Context, string) (*entity.Admin, error)
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.Admin, error)
	}
	Employee interface {
		Create(context.Context, dto.CreateEmployee) (*entity.Employee, error)
		FindOne(context.Context, entity.Employee) (*entity.Employee, error)
		FindAll(context.Context, dto.Page) ([]entity.Employee, error)
		Update(context.Context, uuid.UUID, dto.UpdateEmployee) error
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.Employee, error)
		Delete(context.Context, uuid.UUID) error
	}
	EmployeeRepo interface {
		Create(context.Context, dto.CreateEmployee) (*entity.Employee, error)
		FindOneById(context.Context, uuid.UUID) (*entity.Employee, error)
		FindOneByEmail(context.Context, string) (*entity.Employee, error)
		FindOneByRefreshToken(context.Context, string) (*entity.Employee, error)
		FindAll(context.Context, dto.Page) ([]entity.Employee, error)
		Update(context.Context, uuid.UUID, dto.UpdateEmployee) error
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.Employee, error)
		Delete(context.Context, uuid.UUID) error
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
	Client interface {
		Create(context.Context, dto.CreateClient) (*entity.Client, error)
		Update(context.Context, uuid.UUID, dto.UpdateClient) error
		FindOne(context.Context, entity.Client) (*entity.Client, error)
		FindAll(context.Context, dto.Page) ([]entity.Client, error)
		Delete(context.Context, uuid.UUID) error
	}
	ClientRepo interface {
		Create(context.Context, dto.CreateClient) (*entity.Client, error)
		Update(context.Context, uuid.UUID, dto.UpdateClient) error
		FindOneById(context.Context, uuid.UUID) (*entity.Client, error)
		FindOneByEmail(context.Context, string) (*entity.Client, error)
		FindOneByPhone(context.Context, string) (*entity.Client, error)
		FindAll(context.Context, dto.Page) ([]entity.Client, error)
		Delete(context.Context, uuid.UUID) error
	}
	RealEstate interface {
		Create(context.Context, dto.CreateRealEstate) (*entity.RealEstate, error)
		Update(context.Context, uuid.UUID, dto.UpdateRealEstate) error
		FindOne(context.Context, entity.RealEstate) (*entity.RealEstate, error)
		FindAll(context.Context, dto.Page) ([]entity.RealEstate, error)
		Delete(context.Context, uuid.UUID) error
	}
	RealEstateRepo interface {
		Create(context.Context, dto.CreateRealEstate) (*entity.RealEstate, error)
		Update(context.Context, uuid.UUID, dto.UpdateRealEstate) error
		FindOneById(context.Context, uuid.UUID) (*entity.RealEstate, error)
		FindOneByClientId(context.Context, uuid.UUID) (*entity.RealEstate, error)
		FindAll(context.Context, dto.Page) ([]entity.RealEstate, error)
		Delete(context.Context, uuid.UUID) error
	}
)
