package repo

import (
	"context"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/postgres"
)

type UserRepo struct {
	*postgres.Postgres
}

func NewUser(db *postgres.Postgres) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) Create(ctx context.Context, data dto.CreateUser) (*entity.User, error) {
	return &entity.User{}, nil
}

func (r *UserRepo) FindAll(ctx context.Context) ([]entity.User, error) {
	return []entity.User{}, nil
}

func (r *UserRepo) Delete(ctx context.Context, id string) error {
	return nil
}
