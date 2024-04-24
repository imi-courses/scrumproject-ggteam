package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (r *Repository) FindOneById(ctx context.Context, id uuid.UUID) (*entity.RealEstate, error) {
	var realEstate *entity.RealEstate
	if err := r.WithContext(ctx).Where("id = ?", id).First(&realEstate).Error; err != nil {
		return nil, err
	}
	return realEstate, nil
}

func (r *Repository) FindOneByClientId(
	ctx context.Context,
	id uuid.UUID,
) (*entity.RealEstate, error) {
	var realEstate *entity.RealEstate
	if err := r.WithContext(ctx).Where("client_id = ?", id).First(&realEstate).Error; err != nil {
		return nil, err
	}
	return realEstate, nil
}
