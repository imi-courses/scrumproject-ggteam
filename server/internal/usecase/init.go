package usecase

import (
	repoAdmin "github.com/imi-courses/scrumproject-ggteam/server/internal/repository/admin"
	repoEmployee "github.com/imi-courses/scrumproject-ggteam/server/internal/repository/employee"
	usecaseAdmin "github.com/imi-courses/scrumproject-ggteam/server/internal/usecase/admin"
	usecaseEmployee "github.com/imi-courses/scrumproject-ggteam/server/internal/usecase/employee"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase/hash"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase/jwt"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/config"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/postgres"
)

type UseCases struct {
	AdminUseCase    Admin
	EmployeeUseCase Employee
	HashUseCase     Hash
	JwtUseCase      Jwt
}

func New(cfg *config.Config, db *postgres.Postgres) UseCases {
	t := cfg.ContextTimeout
	return UseCases{
		AdminUseCase:    usecaseAdmin.New(repoAdmin.New(db), t),
		EmployeeUseCase: usecaseEmployee.New(repoEmployee.New(db), t),
		HashUseCase:     hash.New(),
		JwtUseCase:      jwt.New(jwt.Config(cfg.JWT)),
	}
}
