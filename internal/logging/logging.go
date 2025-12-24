package logging

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Options struct {
	Env string // local/dev/prod
}

func Setup(opts Options) zerolog.Logger {
	// Default time format for logs
	zerolog.TimeFieldFormat = time.RFC3339Nano

	// In local/dev, use pretty console logs. In prod, use JSON.
	var logger zerolog.Logger
	if opts.Env == "local" || opts.Env == "dev" {
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339Nano}
		logger = zerolog.New(output).With().Timestamp().Logger()
	} else {
		logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}

	// Set global level (can be made configurable later)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	return logger
}
