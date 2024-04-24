package repository

import (
	"context"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (r *Repository) Create(
	ctx context.Context,
	data dto.CreateRealEstate,
) (*entity.RealEstate, error) {
	realEstate := &entity.RealEstate{
		Address:  data.Address,
		Type:     data.Type,
		ClientID: data.ClientID,
	}
	if err := r.WithContext(ctx).Create(realEstate).Error; err != nil {
		return nil, err
	}
	return realEstate, nil
}
