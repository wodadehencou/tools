package logging

import "testing"

func TestLogging(t *testing.T) {
	logger := GetLogger("test")
	logger.Info().Str("key", "value").Msg("test logging")
	logger.Warn().Int("key2", 10).Msg("test logging warn")
	logger.Error().Int("key2", 10).Msg("test logging error")
}
