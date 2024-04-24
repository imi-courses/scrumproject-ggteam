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
	data dto.UpdateRealEstate,
) error {
	realEstate := &entity.RealEstate{ID: id}
	if err := r.WithContext(ctx).Model(realEstate).Updates(entity.RealEstate{
		Address:  data.Address,
		Type:     data.Type,
		ClientID: data.ClientID,
	}).Error; err != nil {
		return err
	}
	return nil
}
