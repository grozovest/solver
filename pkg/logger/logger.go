package logger

import "go.uber.org/zap"

type Logger struct {
	Logger *zap.Logger
}

func New(debug bool) (*Logger, error) {
	logger, err := checkDebugLevel(debug)

	return &Logger{Logger: logger}, err
}

func checkDebugLevel(debug bool) (*zap.Logger, error) {
	if debug {
		return zap.NewProduction()
	}

	return zap.NewProduction()
}
