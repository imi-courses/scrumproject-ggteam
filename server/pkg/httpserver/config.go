package httpserver

import (
	"time"
)

type Config struct {
	Port            string        `yaml:"port"             env-required:"true"`
	Host            string        `yaml:"host"             env-required:"true"`
	ReadTimeout     time.Duration `yaml:"read_timeout"     env-required:"true"`
	WriteTimeout    time.Duration `yaml:"write_timeout"    env-required:"true"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env-required:"true"`
}
