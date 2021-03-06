package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

const ENV_KEY = "LOGDIR"

var (
	MaxBackups = 5
	MaxSize    = 5  // megabytes
	MaxAge     = 30 // days
)

var globalLogger zerolog.Logger

func AutoLogger() zerolog.Logger {
	logger := zerolog.New(os.Stderr)
	event := logger.With().Timestamp()

	taskID, ok := os.LookupEnv("TASK_ID")
	if ok {
		event = event.Str("task_id", taskID)
	}

	appID, ok := os.LookupEnv("APP_ID")
	if ok {
		event = event.Str("app_id", appID)
	}

	nodeID, ok := os.LookupEnv("NODE_ID")
	if ok {
		event = event.Str("node_id", nodeID)
	}

	event = event.Str("platform_id", "9n-mpc")

	return event.Logger()
	// var writers []io.Writer
	// writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})

	// logDir, ok := os.LookupEnv(ENV_KEY)
	// if ok {
	// 	writers = append(writers,
	// 		newRollingFile(logDir),
	// 		// zerolog.New(os.Stderr),
	// 	)
	// 	// } else {
	// 	// writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	// }

	// mw := io.MultiWriter(writers...)

	// // zerolog.SetGlobalLevel(zerolog.DebugLevel)
	// return zerolog.New(mw).With().Timestamp().Logger()
}

func newRollingFile(dir string) io.Writer {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Error().Err(err).Str("path", dir).Msg("can't create log directory")
		return nil
	}

	now := time.Now().Unix()
	return &lumberjack.Logger{
		Filename:   path.Join(dir, fmt.Sprintf("log-%d.txt", now)),
		MaxBackups: MaxBackups, // files
		MaxSize:    MaxSize,    // megabytes
		MaxAge:     MaxAge,     // days
	}
}
