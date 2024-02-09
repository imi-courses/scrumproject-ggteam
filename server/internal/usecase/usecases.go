package usecase

import (
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase/repo"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/config"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/postgres"
)

type UseCases struct {
	UserUseCase User
}

func New(cfg *config.Config, db *postgres.Postgres) UseCases {
	return UseCases{
		UserUseCase: NewUser(repo.NewUser(db), cfg.ContextTimeout),
	}
}
