package logger

import "go.uber.org/zap"

// Log is a global zap logger
var Log *zap.Logger

// InitLogger initializes zap logger
func InitLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("failed to initialize logger")
	}

	Log = logger
}
