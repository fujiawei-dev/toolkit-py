{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"os"

	"github.com/rs/zerolog"
)

const TimeFormat = "2006-01-02 15:04:05"

// https://github.com/rs/zerolog
var log zerolog.Logger

func init() {
	log = zerolog.New(
		zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: TimeFormat,
			NoColor:    false},
	).
		With().
		Timestamp().
		CallerWithSkipFrameCount(2).
		Logger()

	SetLevel(zerolog.DebugLevel)
}

func Logger() zerolog.Logger {
	return log
}

func SetLevel(l zerolog.Level) {
	zerolog.SetGlobalLevel(l)
}

func SetLogger(logger zerolog.Logger) {
	log = logger
}
