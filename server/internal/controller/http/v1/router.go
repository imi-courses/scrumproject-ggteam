package v1

import (
	"log/slog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/admin"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/auth"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase"
)

func NewRouter(handler *gin.Engine, l *slog.Logger, uc usecase.UseCases) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	handler.Use(cors.Default())

	h := handler.Group("/api/v1")
	{
		admin.New(h, uc.AdminUseCase, uc.HashUseCase, l)
		auth.New(h, uc.AdminUseCase, uc.EmployeeUseCase, uc.HashUseCase, uc.JwtUseCase, l)
	}
}
