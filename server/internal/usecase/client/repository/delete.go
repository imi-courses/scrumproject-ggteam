package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.WithContext(ctx).Delete(&entity.Client{ID: id}).Error; err != nil {
		return err
	}
	return nil
}
