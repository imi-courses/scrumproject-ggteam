package repository

import (
	"context"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (r *Repository) Create(ctx context.Context, data dto.CreateClient) (*entity.Client, error) {
	client := &entity.Client{
		Firstname:  data.Firstname,
		Surname:    data.Surname,
		Middlename: data.Middlename,
		Email:      data.Email,
		Phone:      data.Phone,
	}
	if err := r.WithContext(ctx).Create(client).Error; err != nil {
		return nil, err
	}
	return client, nil
}
