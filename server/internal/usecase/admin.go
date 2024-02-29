package usecase

import (
	"context"
	"time"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type AdminUseCase struct {
	repo       AdminRepo
	ctxTimeout time.Duration
}

func newAdmin(r AdminRepo, t time.Duration) *AdminUseCase {
	return &AdminUseCase{
		repo:       r,
		ctxTimeout: t,
	}
}

func (uc *AdminUseCase) SignUp(c context.Context, data dto.SignUpAdmin) (*entity.Admin, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	user, err := uc.repo.Create(ctx, dto.CreateAdmin(data))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *AdminUseCase) FindOne(c context.Context, p dto.FindOneAdmin) (*entity.Admin, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	user, err := uc.repo.FindOne(ctx, p)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *AdminUseCase) Delete(c context.Context, id string) error {
	return nil
}

func (uc *AdminUseCase) UpdateRefreshToken(c context.Context, data dto.UpdateRefreshToken) (*entity.Admin, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	admin, err := uc.repo.UpdateRefreshToken(ctx, data)
	if err != nil {
		return nil, err
	}
	return admin, err
}
