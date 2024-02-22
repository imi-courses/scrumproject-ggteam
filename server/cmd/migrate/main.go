package main

import (
	"log"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/config"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/postgres"
)

func main() {
	cfg := config.MustLoad()
	db, err := postgres.New((*postgres.Config)(&cfg.DB))
	if err != nil {
		log.Fatalf("Failed connect to postgres: %s", err.Error())
		return
	}
	db.AutoMigrate(&entity.Admin{}, &entity.Employee{})
}
