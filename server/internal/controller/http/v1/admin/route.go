package admin

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase"
)

type route struct {
	ua usecase.Admin
	uh usecase.Hash
	l  *slog.Logger
}

func New(handler *gin.RouterGroup, ua usecase.Admin, uh usecase.Hash, l *slog.Logger) {
	r := &route{ua, uh, l}
	h := handler.Group("/admin")
	{
		h.GET("/", r.findOne)
	}
}
