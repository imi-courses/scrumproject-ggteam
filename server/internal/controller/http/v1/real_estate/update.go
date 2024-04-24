package realestate

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/exception"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type updateRequest dto.UpdateRealEstate

type updateResponse struct {
	Message string `json:"message"`
}

func (r *route) update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		exception.BadRequest(c, err.Error())
		return
	}

	realEstate, err := r.uc.FindOne(c.Request.Context(), entity.RealEstate{ID: id})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			exception.BadRequest(c, err.Error())
			return
		}
		exception.InternalServerError(c, err.Error())
		return
	}

	var body updateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		exception.BadRequest(c, err.Error())
		return
	}

	if err := r.uc.Update(c.Request.Context(), id, dto.UpdateRealEstate(body)); err != nil {
		exception.InternalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, updateResponse{
		Message: fmt.Sprintf(
			"real estate with this client id %s successfully updated",
			realEstate.ClientID,
		),
	})
}
