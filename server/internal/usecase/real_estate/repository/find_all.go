package repository

import (
	"context"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (r *Repository) FindAll(ctx context.Context, page dto.Page) ([]entity.RealEstate, error) {
	var realEstates []entity.RealEstate
	if err := r.WithContext(ctx).Offset(page.Count * (page.CurrentPage - 1)).Limit(page.Count).Find(&realEstates).Error; err != nil {
		return nil, err
	}
	return realEstates, nil
}
