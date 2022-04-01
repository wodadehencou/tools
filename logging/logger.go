package logging

import (
	"sync"

	"github.com/rs/zerolog"
)

var loggers map[string]*zerolog.Logger
var lock sync.Mutex

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.TimestampFieldName = "t"
	zerolog.LevelFieldName = "l"
	zerolog.MessageFieldName = "m"
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
