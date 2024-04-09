package app

import (
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase/admin"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase/employee"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase/hash"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase/jwt"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/config"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/postgres"
)

func NewUseCases(cfg *config.Config, db *postgres.Postgres) usecase.UseCases {
	t := cfg.ContextTimeout
	return usecase.UseCases{
		AdminUseCase:    admin.Init(db, t),
		EmployeeUseCase: employee.Init(db, t),
		HashUseCase:     hash.Init(),
		JwtUseCase:      jwt.Init(jwt.Config(cfg.JWT)),
	}
}
