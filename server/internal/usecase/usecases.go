package usecase

import (
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase/repo"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/config"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/postgres"
)

type UseCases struct {
	AdminUseCase Admin
	HashUseCase  Hash
	JwtUseCase   Jwt
}

func New(cfg *config.Config, db *postgres.Postgres) UseCases {
	return UseCases{
		AdminUseCase: newAdmin(repo.NewAdmin(db), cfg.ContextTimeout),
		HashUseCase:  newHash(),
		JwtUseCase:   newJwt(JwtConfig(cfg.JWT)),
	}
}
