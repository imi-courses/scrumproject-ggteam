package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (r *Repository) Update(
	ctx context.Context,
	id uuid.UUID,
	data dto.UpdateClient,
) error {
	client := &entity.Client{ID: id}
	if err := r.WithContext(ctx).Model(client).Updates(entity.Client{
		Phone:      data.Phone,
		Email:      data.Email,
		Firstname:  data.Firstname,
		Surname:    data.Surname,
		Middlename: data.Middlename,
	}).Error; err != nil {
		return err
	}
	return nil
}
