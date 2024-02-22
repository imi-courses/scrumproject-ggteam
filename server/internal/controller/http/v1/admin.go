package v1

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase"
)

type adminRoute struct {
	u usecase.Admin
	l *slog.Logger
}

func newAdmin(handler *gin.RouterGroup, u usecase.Admin, l *slog.Logger) {
	r := &adminRoute{u, l}
	h := handler.Group("/admin")
	{
		h.POST("/", r.signUp)
	}
}

type signUpAdminRequest dto.CreateAdmin

func (r *adminRoute) signUp(c *gin.Context) {
	var body signUpAdminRequest

	if err := c.BindJSON(&body); err != nil {
		r.l.Error(err.Error())
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	admin, err := r.u.SignUp(c.Request.Context(), dto.SignUpAdmin(body))
	if err != nil {
		r.l.Error(err.Error())
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, admin)
}
