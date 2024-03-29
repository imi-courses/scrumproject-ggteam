package admin

import (
	"context"

	"github.com/google/uuid"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (r *Repository) FindOneById(ctx context.Context, id uuid.UUID) (*entity.Admin, error) {
	var admin *entity.Admin
	if err := r.WithContext(ctx).Where("id = ?", id).First(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *Repository) FindOneByEmail(ctx context.Context, email string) (*entity.Admin, error) {
	var admin *entity.Admin
	if err := r.WithContext(ctx).Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *Repository) FindOneByRefreshToken(ctx context.Context, refreshToken string) (*entity.Admin, error) {
	var admin *entity.Admin
	if err := r.WithContext(ctx).Where("refresh_token = ?", refreshToken).First(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}
