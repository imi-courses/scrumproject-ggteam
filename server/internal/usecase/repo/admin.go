package repo

import (
	"context"

	"github.com/google/uuid"

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

func (r *AdminRepo) FindOne(ctx context.Context, id uuid.UUID) (*entity.Admin, error) {
	var user *entity.Admin
	if err := r.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *AdminRepo) Delete(ctx context.Context, id string) error {
	return nil
}
