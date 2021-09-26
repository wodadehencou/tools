package logging

import (
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var loggers map[string]*zerolog.Logger
var lock sync.RWMutex

func init() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	loggers = make(map[string]*zerolog.Logger)
}

// func NewLogger() *zerolog.Logger {
// 	// logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
// 	logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
// 	return &logger
// }

func SetLogLevel(level zerolog.Level) {
	log.Log().Stringer("set level", level).Msg("set log level")

	lock.Lock()
	defer lock.Unlock()

	for pkg, logger := range loggers {
		log.Trace().Stringer("set level", level).Str("logger", pkg).Msg("set log level of logger")
		*logger = logger.Level(level)
	}
}

func SetParty(party int) {

	lock.Lock()
	defer lock.Unlock()

	for _, logger := range loggers {
		*logger = logger.With().Int("Party", party).Logger()
	}
}

func SetPackageLogLevel(pkg string, level zerolog.Level) {
	log.Log().Stringer("set level", level).Str("logger", pkg).Msg("set log level of logger")

	lock.Lock()
	defer lock.Unlock()

	logger, ok := loggers[pkg]
	if ok {
		*logger = logger.Level(level)
	}
}

func GetLogger(pkg string) *zerolog.Logger {
	// logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Str("pkg", pkg).Logger()
	logger := zerolog.New(os.Stderr).With().Timestamp().Str("pkg", pkg).Logger()
	lock.Lock()
	defer lock.Unlock()
	loggers[pkg] = &logger
	return &logger
}
