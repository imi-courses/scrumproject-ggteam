package employee

import (
	"time"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/repository"
)

type UseCase struct {
	repo       repository.Employee
	ctxTimeout time.Duration
}

func New(r repository.Employee, t time.Duration) *UseCase {
	return &UseCase{
		repo:       r,
		ctxTimeout: t,
	}
}
