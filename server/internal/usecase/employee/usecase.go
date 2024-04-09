package employee

import (
	"time"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase"
)

type UseCase struct {
	repo       usecase.EmployeeRepo
	ctxTimeout time.Duration
}

func New(r usecase.EmployeeRepo, t time.Duration) *UseCase {
	return &UseCase{
		repo:       r,
		ctxTimeout: t,
	}
}
