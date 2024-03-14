package config

import (
	"time"
)

type JWT struct {
	AccessTokenSecret  string        `yaml:"access_token_secret"  env-required:"true"`
	RefreshTokenSecret string        `yaml:"refresh_token_secret" env-required:"true"`
	AccessExpiresAt    time.Duration `yaml:"access_expires_at"    env-required:"true"`
	RefreshExpiresAt   time.Duration `yaml:"refresh_expires_at"   env-required:"true"`
}
