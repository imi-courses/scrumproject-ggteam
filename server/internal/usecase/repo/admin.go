package repo

import (
	"context"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/postgres"
)

type AdminRepo struct {
	*postgres.Postgres
}

func NewAdmin(db *postgres.Postgres) *AdminRepo {
	return &AdminRepo{db}
}

func (r *AdminRepo) Create(ctx context.Context, data dto.CreateAdmin) (*entity.Admin, error) {
	admin := &entity.Admin{
		Email:    data.Email,
		Password: data.Password,
	}
	if err := r.WithContext(ctx).Create(admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *AdminRepo) FindOne(ctx context.Context, p dto.FindOneAdmin) (*entity.Admin, error) {
	var admin *entity.Admin
	if err := r.WithContext(ctx).Where(&entity.Admin{ID: p.ID, Email: p.Email}).First(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *AdminRepo) Delete(ctx context.Context, id string) error {
	return nil
}

func (r *AdminRepo) UpdateRefreshToken(ctx context.Context, data dto.UpdateRefreshToken) (*entity.Admin, error) {
	admin := &entity.Admin{
		ID: data.ID,
	}
	if err := r.WithContext(ctx).Model(&admin).Update("refresh_token", data.RefreshToken).Error; err != nil {
		return nil, err
	}
	if err := r.WithContext(ctx).Where("id = ?", data.ID).First(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}
