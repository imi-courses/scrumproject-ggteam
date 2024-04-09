package admin

import (
	"time"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase"
)

type UseCase struct {
	repo       usecase.AdminRepo
	ctxTimeout time.Duration
}

func New(r usecase.AdminRepo, t time.Duration) *UseCase {
	return &UseCase{
		repo:       r,
		ctxTimeout: t,
	}
}
