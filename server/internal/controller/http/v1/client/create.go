package client

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/exception"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type createRequest dto.CreateClient

type createResponse *entity.Client

func (r *route) create(c *gin.Context) {
	var body createRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		exception.BadRequest(c, err.Error())
		return
	}

	createClient := dto.CreateClient{
		Firstname:  body.Firstname,
		Middlename: body.Middlename,
		Surname:    body.Surname,
		Email:      body.Email,
		Phone:      body.Phone,
	}

	client, err := r.uc.Create(c.Request.Context(), createClient)
	if err != nil {
		r.l.Error(err.Error())
		exception.InternalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, createResponse(client))
}
