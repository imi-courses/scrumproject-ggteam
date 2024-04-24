package realestate

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase"
)

type route struct {
	uc usecase.RealEstate
	l  *slog.Logger
}

func New(handler *gin.RouterGroup, uc usecase.RealEstate, l *slog.Logger) {
	r := &route{uc, l}
	h := handler.Group("/real-estate")
	{
		h.POST("/create", r.create)
		h.GET("/find-all", r.findAll)
		h.PUT("/:id", r.update)
		h.DELETE("/:id", r.delete)
	}
}
