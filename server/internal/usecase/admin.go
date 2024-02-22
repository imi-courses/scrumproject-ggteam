package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type AdminUseCase struct {
	repo       AdminRepo
	hpass      HashPassword
	ctxTimeout time.Duration
}

func NewAdmin(r AdminRepo, h HashPassword, t time.Duration) *AdminUseCase {
	return &AdminUseCase{
		repo:       r,
		hpass:      h,
		ctxTimeout: t,
	}
}

func (uc *AdminUseCase) SignUp(c context.Context, data dto.SignUpAdmin) (*entity.Admin, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	hashedPassword, err := uc.hpass.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}
	data.Password = hashedPassword
	user, err := uc.repo.Create(ctx, dto.CreateAdmin(data))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *AdminUseCase) SignIn(c context.Context, id uuid.UUID) (*entity.Admin, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	user, err := uc.repo.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *AdminUseCase) FindOne(c context.Context, id uuid.UUID) (*entity.Admin, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	user, err := uc.repo.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *AdminUseCase) Delete(c context.Context, id string) error {
	return nil
}
