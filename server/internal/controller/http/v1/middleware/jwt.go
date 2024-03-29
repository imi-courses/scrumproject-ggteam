package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/exception"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase"
)

func JwtCheck(uj usecase.Jwt) gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.Request.Header["Authorization"]
		if len(headerToken) == 0 {
			exception.UnAuthorized(c)
			return
		}
		token := strings.Split(headerToken[0], " ")[1]
		ok, err := uj.IsTokenValid(token, true)
		if err != nil {
			exception.UnAuthorizedWithMessage(c, err.Error())
			return
		}
		if !ok {
			exception.UnAuthorized(c)
			return
		}
		c.Next()
	}
}
