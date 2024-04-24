package app

import (
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase/admin"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase/client"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase/employee"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase/hash"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase/jwt"
	realestate "github.com/imi-courses/scrumproject-ggteam/server/internal/usecase/real_estate"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/config"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/postgres"
)

func NewUseCases(cfg *config.Config, db *postgres.Postgres) usecase.UseCases {
	t := cfg.ContextTimeout
	return usecase.UseCases{
		AdminUseCase:      admin.Init(db, t),
		EmployeeUseCase:   employee.Init(db, t),
		ClientUseCase:     client.Init(db, t),
		RealEstateUseCase: realestate.Init(db, t),
		HashUseCase:       hash.Init(),
		JwtUseCase:        jwt.Init(jwt.Config(cfg.JWT)),
	}
}
