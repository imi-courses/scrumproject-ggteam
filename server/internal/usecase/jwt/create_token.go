package jwt

import (
	"time"

	_jwt "github.com/golang-jwt/jwt/v5"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
)

type CustomClaims struct {
	dto.TokenPayload
	_jwt.RegisteredClaims
}

func (j *UseCase) CreateToken(payload dto.TokenPayload, isAccessToken bool) (string, error) {
	var expiresAt time.Duration
	if isAccessToken {
		expiresAt = j.AccessExpiresAt
	} else {
		expiresAt = j.RefreshExpiresAt
	}
	claims := CustomClaims{
		payload,
		_jwt.RegisteredClaims{
			ExpiresAt: _jwt.NewNumericDate(time.Now().Add(expiresAt)),
		},
	}
	token := _jwt.NewWithClaims(_jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return t, err
}
