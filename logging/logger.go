package logging

import (
	"sync"

	"github.com/rs/zerolog"
)

var loggers map[string]*zerolog.Logger
var lock sync.Mutex

func init() {
	// zerolog.TimestampFieldName = "time"
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05.999999"

	// zerolog.LevelFieldName = "level"
	// zerolog.LevelDebugValue = "debug"
	// zerolog.LevelInfoValue = "info"
	// zerolog.LevelWarnValue = "warn"
	// zerolog.LevelErrorValue = "error"
	// zerolog.LevelPanicValue = "panic"
	// zerolog.LevelFatalValue = "fatal"

	zerolog.MessageFieldName = "msg"

	globalLogger = AutoLogger()
	loggers = make(map[string]*zerolog.Logger, 32)
}

func SetLogLevel(level zerolog.Level) {
	lock.Lock()
	defer lock.Unlock()
	for _, lp := range loggers {
		*lp = (*lp).Level(level)
	}
}

func SetPackageLevel(pkg string, level zerolog.Level) {
	lock.Lock()
	defer lock.Unlock()
	lp, ok := loggers[pkg]
	if ok {
		*lp = (*lp).Level(level)
	}
}

func GetLogger(pkg string) *zerolog.Logger {
	logger := globalLogger.With().Str("pkg", pkg).Logger()
	lp := &logger
	lock.Lock()
	defer lock.Unlock()
	loggers[pkg] = lp
	return lp
}

func SetLoggerField(ss ...string) {
	if len(ss)%2 != 0 {
		globalLogger.Error().Msg("not any pairs of key-value")
		return
	}
	if len(ss) == 0 {
		return
	}

	lock.Lock()
	defer lock.Unlock()
	for _, lp := range loggers {
		e := (*lp).With()
		for i := 0; i < len(ss); i += 2 {
			e = e.Str(ss[i], ss[i+1])
		}
		*lp = e.Logger()
	}
}
