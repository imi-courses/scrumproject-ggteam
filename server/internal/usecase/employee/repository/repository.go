package repository

import "github.com/imi-courses/scrumproject-ggteam/server/pkg/postgres"

type Repository struct {
	*postgres.Postgres
}

func New(db *postgres.Postgres) *Repository {
	return &Repository{db}
}
