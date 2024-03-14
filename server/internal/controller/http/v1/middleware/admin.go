package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/exception"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/usecase"
)

func AdminCheck(uj usecase.Jwt, ua usecase.Admin) gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.Request.Header["Authorization"]
		if len(headerToken) == 0 {
			exception.UnAuthorized(c)
			return
		}
		token := strings.Split(headerToken[0], " ")[1]
		ok, err := uj.IsTokenValid(token, true)
		if err != nil {
			exception.UnAuthorized(c)
			return
		}
		if !ok {
			exception.UnAuthorized(c)
			return
		}

		id, err := uj.ExtractFromToken(token, "id", true)
		if err != nil {
			exception.UnAuthorized(c)
			return
		}
		email, err := uj.ExtractFromToken(token, "email", true)

		admin, err := ua.FindOne(c, dto.FindOneAdmin{
			ID: uuid.MustParse(id),
		})
		if err != nil {
			exception.UnAuthorized(c)
			return
		}

		if admin.Email != email {
			exception.UnAuthorized(c)
			return
		}

		fmt.Println(id)

		c.Next()
	}
}
