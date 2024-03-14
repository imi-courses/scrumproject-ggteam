package admin

import (
	"context"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (r *Repository) FindOne(ctx context.Context, p dto.FindOneAdmin) (*entity.Admin, error) {
	var admin *entity.Admin
	if err := r.WithContext(ctx).Where(&entity.Admin{ID: p.ID, Email: p.Email}).First(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}
