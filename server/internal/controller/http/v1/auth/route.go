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

func New(handler *gin.RouterGroup, ua usecase.Admin, ue usecase.Employee, uh usecase.Hash, uj usecase.Jwt, l *slog.Logger) {
	r := &route{ua, ue, uh, uj, l}
	h := handler.Group("/auth")
	{
		h.POST("/admin", r.signInAdmin)
		h.POST("/", r.singUp)
	}
}
