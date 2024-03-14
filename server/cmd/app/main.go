package main

import (
	"github.com/imi-courses/scrumproject-ggteam/server/internal/app"
	"github.com/imi-courses/scrumproject-ggteam/server/pkg/config"
)

func main() {
	cfg := config.MustLoad()

	app.Run(cfg)
}
