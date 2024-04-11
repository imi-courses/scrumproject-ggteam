package employee

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase"
)

type route struct {
	ue usecase.Employee
	uh usecase.Hash
	l  *slog.Logger
}

func New(handler *gin.RouterGroup, ue usecase.Employee, uh usecase.Hash, l *slog.Logger) {
	r := &route{ue, uh, l}
	h := handler.Group("/employee")
	{
		h.POST("/create", r.create)
		h.GET("/find-all", r.findAll)
		h.PUT("/:id", r.update)
		h.DELETE("/:id", r.delete)
	}
}
