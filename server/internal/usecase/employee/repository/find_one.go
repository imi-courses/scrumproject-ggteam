package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (r *Repository) FindOneById(ctx context.Context, id uuid.UUID) (*entity.Employee, error) {
	var employee *entity.Employee
	if err := r.WithContext(ctx).Where("id = ?", id).First(&employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}

func (r *Repository) FindOneByEmail(ctx context.Context, email string) (*entity.Employee, error) {
	var employee *entity.Employee
	if err := r.WithContext(ctx).Where("email = ?", email).First(&employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}

func (r *Repository) FindOneByRefreshToken(
	ctx context.Context,
	refreshToken string,
) (*entity.Employee, error) {
	var employee *entity.Employee
	if err := r.WithContext(ctx).Where("refresh_token = ?", refreshToken).First(&employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}
