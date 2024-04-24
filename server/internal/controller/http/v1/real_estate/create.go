package realestate

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/exception"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type createRequest dto.CreateRealEstateRequest

type createResponse *entity.RealEstate

func (r *route) create(c *gin.Context) {
	var body createRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		exception.BadRequest(c, err.Error())
		return
	}

	id, err := uuid.Parse(body.ClientID)
	if err != nil {
		exception.BadRequest(c, err.Error())
		return
	}

	createRealEstate := dto.CreateRealEstate{
		Address:  body.Address,
		Type:     body.Type,
		ClientID: id,
	}

	realEstate, err := r.uc.Create(c.Request.Context(), createRealEstate)
	if err != nil {
		r.l.Error(err.Error())
		exception.InternalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, createResponse(realEstate))
}
