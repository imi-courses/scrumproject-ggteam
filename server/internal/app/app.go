package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	v1 "github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/config"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/httpserver"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/logger"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/postgres"
)

func Run(cfg *config.Config) {
	log := logger.New(cfg.Env)

	log.Info(fmt.Sprintf("Starting server at Port: %s", cfg.HTTP.Port))

	db, err := postgres.New(&cfg.DB)
	if err != nil {
		logger.Fatal(log, "Failed connect to postgres:", err)
	}
	log.Info("Connected to postgres")

	// Migration in prod
	if cfg.Env == "prod" {
		if (!db.Migrator().HasTable(&entity.Admin{})) {
			migrate(db)
		}
	}

	// UseCases
	usecases := usecase.New(cfg, db)

	// HTTP SERVER
	gin.SetMode(gin.ReleaseMode)
	if cfg.Env == "local" {
		gin.SetMode(gin.DebugMode)
	}
	handler := gin.New()
	v1.NewRouter(handler, log, usecases)
	httpServer := httpserver.New(handler, &cfg.HTTP)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		logger.Fatal(log, "app - Run - httpServer.Notify", err)
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		logger.Fatal(log, "app - Run - httpServer.Shutdown", err)
	}
}
