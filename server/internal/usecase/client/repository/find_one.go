package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (r *Repository) FindOneById(ctx context.Context, id uuid.UUID) (*entity.Client, error) {
	var client *entity.Client
	if err := r.WithContext(ctx).Where("id = ?", id).First(&client).Error; err != nil {
		return nil, err
	}
	return client, nil
}

func (r *Repository) FindOneByEmail(ctx context.Context, email string) (*entity.Client, error) {
	var client *entity.Client
	if err := r.WithContext(ctx).Where("email = ?", email).First(&client).Error; err != nil {
		return nil, err
	}
	return client, nil
}

func (r *Repository) FindOneByPhone(
	ctx context.Context,
	phone string,
) (*entity.Client, error) {
	var client *entity.Client
	if err := r.WithContext(ctx).Where("phone = ?", phone).First(&client).Error; err != nil {
		return nil, err
	}
	return client, nil
}
