package client

import (
	"time"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase"
)

type UseCase struct {
	repo       usecase.ClientRepo
	ctxTimeout time.Duration
}

func New(r usecase.ClientRepo, t time.Duration) *UseCase {
	return &UseCase{
		repo:       r,
		ctxTimeout: t,
	}
}
