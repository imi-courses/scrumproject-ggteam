package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/exception"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type meAdminResponse struct {
	Admin *entity.Admin `json:"admin"`
	Role  string        `json:"role"`
}

type meEmployeeResponse struct {
	Employee *entity.Employee `json:"employee"`
	Role     string           `json:"role"`
}

func (r *route) me(c *gin.Context) {
	headerToken := c.Request.Header["Authorization"]
	token := strings.Split(headerToken[0], " ")[1]
	role, err := r.uj.ExtractFromToken(token, "role", true)
	if err != nil {
		exception.UnAuthorized(c)
		return
	}
	id, err := r.uj.ExtractFromToken(token, "id", true)
	if err != nil {
		exception.UnAuthorized(c)
		return
	}
	if role == "admin" {
		admin, err := r.ua.FindOne(c.Request.Context(), entity.Admin{ID: uuid.MustParse(id)})
		if err != nil {
			exception.InternalServerError(c, err.Error())
			return
		}

		c.JSON(http.StatusOK, meAdminResponse{
			Admin: admin,
			Role:  role,
		})
	} else if role == "employee" {
		employee, err := r.ue.FindOne(c.Request.Context(), entity.Employee{ID: uuid.MustParse(id)})
		if err != nil {
			exception.InternalServerError(c, err.Error())
			return
		}

		c.JSON(http.StatusOK, meEmployeeResponse{
			Employee: employee,
			Role:     role,
		})
	}
}
