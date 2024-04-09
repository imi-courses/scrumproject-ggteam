package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/exception"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type signInAdminRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signInAdminResponse struct {
	Admin       *entity.Admin `json:"admin"`
	AccessToken string        `json:"access_token"`
}

func (r *route) signInAdmin(c *gin.Context) {
	var body signInAdminRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		exception.BadRequest(c, err.Error())
		return
	}

	candidate, err := r.ua.FindOne(c, entity.Admin{Email: body.Email})
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
		Role:  "admin",
	}, dto.RefreshTokenPayload{
		ID:    fmt.Sprintf("%v", candidate.ID),
		Email: candidate.Email,
		Role:  "admin",
	})
	if err != nil {
		exception.InternalServerError(c, err.Error())
		return
	}

	admin, err := r.ua.UpdateRefreshToken(c.Request.Context(), dto.UpdateRefreshToken{
		ID:           candidate.ID,
		RefreshToken: tokens.RefreshToken,
	})
	if err != nil {
		exception.InternalServerError(c, err.Error())
		return
	}

	c.SetCookie("refresh_token", tokens.RefreshToken, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, &signInAdminResponse{
		Admin:       admin,
		AccessToken: tokens.AccessToken,
	})
}
