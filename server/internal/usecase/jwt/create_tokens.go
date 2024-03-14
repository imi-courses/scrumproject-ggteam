package jwt

import "github.com/imi-courses/scrumproject-ggteam/server/internal/dto"

func (j *UseCase) CreateTokens(atPayload dto.AccessTokenPayload, rtPayload dto.RefreshTokenPayload) (*dto.Tokens, error) {
	accessToken, err := j.CreateAccessToken(atPayload)
	if err != nil {
		return nil, err
	}
	refreshToken, err := j.CreateRefreshToken(rtPayload)
	if err != nil {
		return nil, err
	}
	return &dto.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
