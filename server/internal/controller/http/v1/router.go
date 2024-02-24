package v1

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase"
)

func NewRouter(handler *gin.Engine, l *slog.Logger, uc usecase.UseCases) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	h := handler.Group("/api/v1")
	{
		newAdmin(h, uc.AdminUseCase, uc.HashUseCase, l)
	}
}
