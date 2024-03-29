package admin

import (
	"time"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase/admin/repository"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/postgres"
)

func Init(db *postgres.Postgres, t time.Duration) *UseCase {
	return New(repository.New(db), t)
}
