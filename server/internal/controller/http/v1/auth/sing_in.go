package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/exception"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type signInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signInResponse struct {
	Employee    *entity.Employee `json:"employee"`
	AccessToken string           `json:"access_token"`
}

func (r *route) signIn(c *gin.Context) {
	var body signInRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		exception.BadRequest(c, err.Error())
		return
	}

	candidate, err := r.ue.FindOne(c, entity.Employee{Email: body.Email})
	if err != nil {
		exception.BadRequest(c, err.Error())
		return
	}

	ok := r.uh.CheckPasswordHash(body.Password, candidate.Password)
	if !ok {
		exception.BadRequest(c, "incorrect password")
		return
	}

	tokens, err := r.uj.CreateTokens(dto.AccessTokenPayload{
		ID:    fmt.Sprintf("%v", candidate.ID),
		Email: candidate.Email,
		Role:  "employee",
	}, dto.RefreshTokenPayload{
		ID:    fmt.Sprintf("%v", candidate.ID),
		Email: candidate.Email,
		Role:  "employee",
	})
	if err != nil {
		exception.InternalServerError(c, err.Error())
		return
	}

	employee, err := r.ue.UpdateRefreshToken(c.Request.Context(), dto.UpdateRefreshToken{
		ID:           candidate.ID,
		RefreshToken: tokens.RefreshToken,
	})
	if err != nil {
		exception.InternalServerError(c, err.Error())
		return
	}

	c.SetCookie("refresh_token", tokens.RefreshToken, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, &signInResponse{
		Employee:    employee,
		AccessToken: tokens.AccessToken,
	})
}
