package client

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase"
)

type route struct {
	uc usecase.Client
	l  *slog.Logger
}

func New(handler *gin.RouterGroup, uc usecase.Client, l *slog.Logger) {
	r := &route{uc, l}
	h := handler.Group("/client")
	{
		h.POST("/", r.create)
		h.GET("/", r.findAll)
	}
}
