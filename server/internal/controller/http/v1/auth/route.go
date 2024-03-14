package auth

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase"
)

type route struct {
	ua usecase.Admin
	ue usecase.Employee
	uh usecase.Hash
	uj usecase.Jwt
	l  *slog.Logger
}

func New(publicHandler *gin.RouterGroup, privateHandler *gin.RouterGroup, ua usecase.Admin, ue usecase.Employee, uh usecase.Hash, uj usecase.Jwt, l *slog.Logger) {
	r := &route{ua, ue, uh, uj, l}
	public := publicHandler.Group("/auth")
	private := privateHandler.Group("/auth")
	{
		public.POST("/admin", r.signInAdmin)
		public.POST("/", r.signIn)
		private.GET("/me", r.me)
		public.GET("/refresh", r.refresh)
	}
}
