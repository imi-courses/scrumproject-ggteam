package v1

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

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
		h.GET("/", r.findOne)
	}
}

// Sign Up

type signUpAdminRequest dto.CreateAdmin

func (r *adminRoute) signUp(c *gin.Context) {
	var body signUpAdminRequest

	if err := c.BindJSON(&body); err != nil {
		r.l.Error(err.Error())
		badRequest(c, err.Error())
		return
	}

	hashedPassword, err := r.uh.HashPassword(body.Password)
	if err != nil {
		r.l.Error(err.Error())
		badRequest(c, err.Error())
		return
	}

	body.Password = hashedPassword

	admin, err := r.ua.SignUp(c.Request.Context(), dto.SignUpAdmin(body))
	if err != nil {
		r.l.Error(err.Error())
		internalServerError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, admin)
}

// Find One

func (r *adminRoute) findOne(c *gin.Context) {
	var id uuid.UUID
	if c.Query("id") == "" && c.Query("email") == "" {
		badRequest(c, "no data")
		return
	}
	if c.Query("id") != "" {
		var err error
		id, err = uuid.Parse(c.Query("id"))
		if err != nil {
			badRequest(c, err.Error())
			return
		}
	}
	p := dto.FindOneAdmin{
		ID:    id,
		Email: c.Query("email"),
	}

	user, err := r.ua.FindOne(c.Request.Context(), p)
	if err != nil {
		r.l.Error(err.Error())
		internalServerError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
