package app

import (
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/postgres"
)

func migrate(db *postgres.Postgres) {
	db.AutoMigrate(&entity.Admin{}, &entity.Employee{}, &entity.Client{}, &entity.RealEstate{})
}
