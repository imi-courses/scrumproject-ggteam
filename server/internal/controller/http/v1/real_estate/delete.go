package realestate

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/exception"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type deleteResponse struct {
	Message string `json:"message"`
}

func (r *route) delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		exception.BadRequest(c, err.Error())
		return
	}

	_, err = r.uc.FindOne(c.Request.Context(), entity.RealEstate{
		ID: id,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			exception.BadRequest(c, err.Error())
		} else {
			exception.InternalServerError(c, err.Error())
		}
		return
	}

	if err := r.uc.Delete(c.Request.Context(), id); err != nil {
		exception.InternalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, deleteResponse{
		Message: "real estate was deleted",
	})
}
