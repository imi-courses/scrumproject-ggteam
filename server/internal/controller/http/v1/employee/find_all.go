package employee

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/exception"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type findAllResponse struct {
	Employees []entity.Employee `json:"employees"`
}

func (r *route) findAll(c *gin.Context) {
	var currentPage, count int
	var err error
	if c.Query("page") == "" {
		currentPage = 1
	} else {
		currentPage, err = strconv.Atoi(c.Query("page"))
		if err != nil {
			exception.BadRequest(c, "page undefined")
			return
		}
	}
	if c.Query("count") == "" {
		count = 10
	} else {
		count, err = strconv.Atoi(c.Query("count"))
		if err != nil {
			exception.BadRequest(c, "count undefined")
			return
		}
	}

	employees, err := r.ue.FindAll(c.Request.Context(), dto.Page{
		Count:       count,
		CurrentPage: currentPage,
	})
	if err != nil {
		exception.InternalServerError(c, "employees not found")
		return
	}

	c.JSON(http.StatusOK, findAllResponse{Employees: employees})
}
