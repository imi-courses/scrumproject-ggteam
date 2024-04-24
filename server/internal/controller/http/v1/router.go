package v1

import (
	"log/slog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/admin"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/auth"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/client"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/employee"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/middleware"
	realestate "github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/real_estate"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase"
)

func NewRouter(handler *gin.Engine, l *slog.Logger, uc usecase.UseCases) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "http://y-student.ru"}
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	handler.Use(cors.New(config))

	public := handler.Group("/api/v1")
	private := handler.Group("/api/v1")
	protected := handler.Group("/api/v1")

	private.Use(middleware.JwtCheck(uc.JwtUseCase))
	protected.Use(middleware.AdminCheck(uc.JwtUseCase, uc.AdminUseCase))
	{
		admin.New(public, uc.AdminUseCase, uc.HashUseCase, l)
		auth.New(
			public,
			private,
			uc.AdminUseCase,
			uc.EmployeeUseCase,
			uc.HashUseCase,
			uc.JwtUseCase,
			l,
		)
		employee.New(protected, uc.EmployeeUseCase, uc.HashUseCase, l)
		client.New(private, uc.ClientUseCase, l)
		realestate.New(private, uc.RealEstateUseCase, l)
	}
}
