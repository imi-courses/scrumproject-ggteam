package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/exception"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type signUpRequest dto.CreateEmployee

type signUpResponse struct {
	Employee     *entity.Employee `json:"employee"`
	AccessToken  string           `json:"access_token"`
	RefreshToken string           `json:"refresh_token"`
}

func (r *route) singUp(c *gin.Context) {
	var body signUpRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		exception.BadRequest(c, err.Error())
		return
	}

	_, err := r.ue.FindOne(c, dto.FindOneEmployee{Email: body.Email})
	if err == nil {
		exception.BadRequest(
			c,
			fmt.Sprintf("A user with this email %s already exist", body.Email),
		)
		return
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		exception.BadRequest(c, err.Error())
		return
	}

	hashedPassword, err := r.uh.HashPassword(body.Password)
	if err != nil {
		exception.InternalServerError(c, err.Error())
		return
	}

	body.Password = hashedPassword

	employee, err := r.ue.SignUp(c, dto.CreateEmployee(body))
	if err != nil {
		exception.InternalServerError(c, err.Error())
		return
	}

	accessToken, err := r.uj.CreateToken(dto.TokenPayload{
		ID:    fmt.Sprintf("%v", employee.ID),
		Email: employee.Email,
	}, true)
	if err != nil {
		exception.InternalServerError(c, err.Error())
		return
	}

	refreshToken, err := r.uj.CreateToken(dto.TokenPayload{
		ID:    fmt.Sprintf("%v", employee.ID),
		Email: employee.Email,
	}, false)
	if err != nil {
		exception.InternalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, &signUpResponse{
		Employee:     employee,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
