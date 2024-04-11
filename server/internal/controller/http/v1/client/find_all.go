package client

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/exception"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/util"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type findAllResponse struct {
	Clients []entity.Client `json:"clients"`
	Page    int             `json:"page"`
	Count   int             `json:"count"`
}

func (r *route) findAll(c *gin.Context) {
	page, count, err := util.GetPage(c.Query("page"), c.Query("count"))
	if err != nil {
		exception.BadRequest(c, err.Error())
		return
	}

	clients, err := r.uc.FindAll(c.Request.Context(), dto.Page{
		Count:       count,
		CurrentPage: page,
	})
	if err != nil {
		exception.InternalServerError(c, "clients not found")
		return
	}

	c.JSON(http.StatusOK, findAllResponse{Clients: clients, Page: page, Count: count})
}
