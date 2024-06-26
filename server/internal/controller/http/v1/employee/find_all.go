package employee

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/exception"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/util"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type findAllResponse struct {
	Employees []entity.Employee `json:"employees"`
	Page      int               `json:"page"`
	Count     int               `json:"count"`
}

func (r *route) findAll(c *gin.Context) {
	var err error
	page, count, err := util.GetPage(c.Query("page"), c.Query("count"))
	if err != nil {
		exception.BadRequest(c, err.Error())
		return
	}

	employees, err := r.ue.FindAll(c.Request.Context(), dto.Page{
		Count:       count,
		CurrentPage: page,
	})
	if err != nil {
		exception.InternalServerError(c, "employees not found")
		return
	}

	c.JSON(http.StatusOK, findAllResponse{Employees: employees, Page: page, Count: count})
}
