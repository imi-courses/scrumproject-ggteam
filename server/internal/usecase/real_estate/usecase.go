package realestate

import (
	"time"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase"
)

type UseCase struct {
	repo       usecase.RealEstateRepo
	ctxTimeout time.Duration
}

func New(r usecase.RealEstateRepo, t time.Duration) *UseCase {
	return &UseCase{
		repo:       r,
		ctxTimeout: t,
	}
}
