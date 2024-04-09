package repository

import (
	"context"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (r *Repository) UpdateRefreshToken(
	ctx context.Context,
	data dto.UpdateRefreshToken,
) (*entity.Employee, error) {
	employee := &entity.Employee{
		ID: data.ID,
	}
	if err := r.WithContext(ctx).Model(&employee).Update("refresh_token", data.RefreshToken).Error; err != nil {
		return nil, err
	}
	if err := r.WithContext(ctx).Where("id = ?", data.ID).First(&employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}
