package logging

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func Setup() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}
}

func Info(msg string, v ...interface{}) {
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Timestamp().Logger()
	log.Info().Msgf(msg, v...)
}

func Error(msg string, v ...interface{}) {
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Timestamp().Logger()
	log.Error().Msgf(msg, v...)
}

func Debug(msg string, v ...interface{}) {
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Timestamp().Logger()
	log.Error().Msgf(msg, v...)
}
