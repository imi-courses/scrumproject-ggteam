package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
)

type CustomAccessTokenClaims struct {
	dto.AccessTokenPayload
	jwt.RegisteredClaims
}

func (j *UseCase) CreateAccessToken(payload dto.AccessTokenPayload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomAccessTokenClaims{
		payload,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.AccessExpiresAt)),
		},
	})
	t, err := token.SignedString([]byte(j.AccessTokenSecret))
	if err != nil {
		return "", err
	}
	return t, err
}
