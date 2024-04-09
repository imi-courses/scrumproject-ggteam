package repository

import (
	"context"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (r *Repository) FindAll(ctx context.Context, page dto.Page) ([]entity.Employee, error) {
	var employees []entity.Employee
	if err := r.WithContext(ctx).Offset(page.Count * (page.CurrentPage - 1)).Limit(page.Count).Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}
