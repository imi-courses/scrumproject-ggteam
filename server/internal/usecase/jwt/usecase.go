package jwt

import (
	"time"
)

type Config struct {
	Secret           string
	AccessExpiresAt  time.Duration
	RefreshExpiresAt time.Duration
}

type UseCase struct {
	Config
}

func New(c Config) *UseCase {
	return &UseCase{c}
}
