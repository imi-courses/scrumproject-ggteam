package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/exception"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (r *route) findOne(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		exception.BadRequest(c, err.Error())
		return
	}

	data := entity.Client{
		ID: id,
	}

	client, err := r.uc.FindOne(c.Request.Context(), data)
	if err != nil {
		r.l.Error(err.Error())
		exception.InternalServerError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, client)
}
