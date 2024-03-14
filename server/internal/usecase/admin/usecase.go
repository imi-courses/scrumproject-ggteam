package admin

import (
	"time"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/repository"
)

type UseCase struct {
	repo       repository.Admin
	ctxTimeout time.Duration
}

func New(r repository.Admin, t time.Duration) *UseCase {
	return &UseCase{
		repo:       r,
		ctxTimeout: t,
	}
}
