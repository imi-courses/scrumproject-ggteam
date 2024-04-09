package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/exception"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type refreshTokensResponse *dto.Tokens

func (r *route) refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		exception.UnAuthorizedWithMessage(c, err.Error())
		return
	}
	if refreshToken == "" {
		exception.UnAuthorizedWithMessage(c, "refresh token is empty")
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

		tokens, err := r.uj.CreateTokens(dto.AccessTokenPayload{
			ID:    fmt.Sprintf("%v", admin.ID),
			Email: admin.Email,
			Role:  role,
		}, dto.RefreshTokenPayload{
			ID:    fmt.Sprintf("%v", admin.ID),
			Email: admin.Email,
			Role:  role,
		})

		_, err = r.ua.UpdateRefreshToken(c, dto.UpdateRefreshToken{ID: admin.ID, RefreshToken: tokens.RefreshToken})

		c.SetCookie("refresh_token", tokens.RefreshToken, 3600, "/", "localhost", false, true)

		if err != nil {
			return
		}

		c.JSON(http.StatusOK, refreshTokensResponse(tokens))
	} else if role == "employee" {
		employee, err := r.ue.FindOne(c, entity.Employee{
			RefreshToken: refreshToken,
		})
		if err != nil {
			exception.UnAuthorized(c)
			return
		}

		tokens, err := r.uj.CreateTokens(dto.AccessTokenPayload{
			ID:    fmt.Sprintf("%v", employee.ID),
			Email: employee.Email,
			Role:  role,
		}, dto.RefreshTokenPayload{
			ID:    fmt.Sprintf("%v", employee.ID),
			Email: employee.Email,
			Role:  role,
		})

		_, err = r.ue.UpdateRefreshToken(c, dto.UpdateRefreshToken{ID: employee.ID, RefreshToken: tokens.RefreshToken})

		c.SetCookie("refresh_token", tokens.RefreshToken, 3600, "/", "localhost", false, true)

		if err != nil {
			return
		}

		c.JSON(http.StatusOK, refreshTokensResponse(tokens))
	}
}
