package logging

import (
	"sync"

	"github.com/rs/zerolog"
)

var loggers []*zerolog.Logger
var lock sync.Mutex

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.TimestampFieldName = "t"
	zerolog.LevelFieldName = "l"
	zerolog.MessageFieldName = "m"
	globalLogger = AutoLogger()
	loggers = make([]*zerolog.Logger, 0, 32)
}

func SetLogLevel(level zerolog.Level) {
	lock.Lock()
	defer lock.Unlock()
	for _, lp := range loggers {
		*lp = (*lp).Level(level)
	}
}

func GetLogger(pkg string) *zerolog.Logger {
	logger := globalLogger.With().Str("pkg", pkg).Logger()
	lp := &logger
	lock.Lock()
	defer lock.Unlock()
	loggers = append(loggers, lp)
	return lp
}
