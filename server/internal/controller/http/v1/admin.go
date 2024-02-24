package v1

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase"
)

type adminRoute struct {
	ua usecase.Admin
	uh usecase.Hash
	l  *slog.Logger
}

func newAdmin(handler *gin.RouterGroup, ua usecase.Admin, uh usecase.Hash, l *slog.Logger) {
	r := &adminRoute{ua, uh, l}
	h := handler.Group("/admin")
	{
		h.POST("/", r.signUp)
	}
}

// Sign Up

type signUpAdminRequest dto.CreateAdmin

func (r *adminRoute) signUp(c *gin.Context) {
	var body signUpAdminRequest

	if err := c.BindJSON(&body); err != nil {
		r.l.Error(err.Error())
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := r.uh.HashPassword(body.Password)
	if err != nil {
		r.l.Error(err.Error())
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	body.Password = hashedPassword

	admin, err := r.ua.SignUp(c.Request.Context(), dto.SignUpAdmin(body))
	if err != nil {
		r.l.Error(err.Error())
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, admin)
}
