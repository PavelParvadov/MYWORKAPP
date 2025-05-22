package logger

import (
	"MYWORKAPP/config"
	"github.com/rs/zerolog"
	"os"
)

func NewLogger(config *config.LogConfig) *zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.Level(config.Level))
	var logger zerolog.Logger

	if config.Format == "JSON" {
		logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	} else {
		ConsoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		logger = zerolog.New(ConsoleWriter).With().Timestamp().Logger()
	}
	return &logger
}
