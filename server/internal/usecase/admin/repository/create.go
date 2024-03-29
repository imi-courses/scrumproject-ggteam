package repository

import (
	"context"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (r *Repository) Create(ctx context.Context, data dto.CreateAdmin) (*entity.Admin, error) {
	admin := &entity.Admin{
		Email:    data.Email,
		Password: data.Password,
	}
	if err := r.WithContext(ctx).Create(admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}
