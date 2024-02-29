package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/exception"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
)

func (r *route) findOne(c *gin.Context) {
	var id uuid.UUID
	if c.Query("id") == "" && c.Query("email") == "" {
		exception.BadRequest(c, "no data")
		return
	}
	if c.Query("id") != "" {
		var err error
		id, err = uuid.Parse(c.Query("id"))
		if err != nil {
			exception.BadRequest(c, err.Error())
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
		exception.InternalServerError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
