package config

import (
	"time"
)

type JWT struct {
	Secret           string        `yaml:"secret"             env-required:"true"`
	AccessExpiresAt  time.Duration `yaml:"access_expires_at"  env-required:"true"`
	RefreshExpiresAt time.Duration `yaml:"refresh_expires_at" env-required:"true"`
}
