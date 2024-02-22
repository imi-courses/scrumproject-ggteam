package usecase

import (
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase/repo"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/config"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/postgres"
)

type UseCases struct {
	AdminUseCase Admin
}

func New(cfg *config.Config, db *postgres.Postgres) UseCases {
	return UseCases{
		AdminUseCase: NewAdmin(repo.NewAdmin(db), newHashPassword(), cfg.ContextTimeout),
	}
}
