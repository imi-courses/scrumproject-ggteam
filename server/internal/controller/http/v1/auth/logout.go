package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/exception"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type logoutResponse struct {
	Message string `json:"message"`
}

func (r *route) logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		exception.BadRequest(c, "cookie not set")
		return
	}
	if refreshToken == "" {
		exception.BadRequest(c, "refresh token is empty")
		return
	}

	role, err := r.uj.ExtractFromToken(refreshToken, "role", false)
	if err != nil {
		exception.UnAuthorized(c)
		return
	}

	if role == "admin" {
		admin, err := r.ua.FindOne(c, entity.Admin{
			RefreshToken: refreshToken,
		})
		if err != nil {
			exception.UnAuthorized(c)
			return
		}

		_, err = r.ua.UpdateRefreshToken(c, dto.UpdateRefreshToken{ID: admin.ID, RefreshToken: ""})
		if err != nil {
			exception.InternalServerError(c, "can not delete refresh token in db")
			return
		}
	} else if role == "employee" {
		user, err := r.ue.FindOne(c, entity.Employee{
			RefreshToken: refreshToken,
		})
		if err != nil {
			exception.UnAuthorized(c)
			return
		}

		_, err = r.ue.UpdateRefreshToken(c, dto.UpdateRefreshToken{ID: user.ID, RefreshToken: ""})
		if err != nil {
			exception.InternalServerError(c, "can not delete refresh token in db")
			return
		}
	} else {
		exception.InternalServerError(c, fmt.Sprintf("this role %s is not exists", role))
		return
	}
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, logoutResponse{Message: "successfully logged out"})
}
