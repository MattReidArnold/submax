package main

import (
	"log/slog"
	"os"

	"github.com/caarlos0/env"
	"github.com/mattreidarnold/submax/internal/service"
	"github.com/pkg/errors"
)

type config struct {
	Port        int    `env:"PORT" envDefault:"8000"`
	ContextPath string `env:"CONTEXT_PATH" envDefault:"/v1"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		panic(errors.Wrap(err, "parsing env config"))
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	svc := service.NewService(logger, cfg.Port, cfg.ContextPath)

	err := svc.Run()
	if err != nil {
		panic(errors.Wrap(err, "running service"))
	}
}
