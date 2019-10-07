package logger

import (
	"go.uber.org/zap"
)

// https://github.com/uber-go/zap
// We use zap to log,if someone wants to replace it with another logging lib, change it here
// note that we have to import "go.uber.org/zap" instead of github

// Logger is a interface for leveled logging
// We will use template logging instead of structured logging
type Logger interface {
	Log(template string, args ...interface{})
	Debug(template string, args ...interface{})
	Info(template string, args ...interface{})
	Warn(template string, args ...interface{})
	Error(template string, args ...interface{})
}

// LoggingService is a struct implement logger interface
type LoggingService struct {
	logger *zap.SugaredLogger
}

// NewLogger return a logger instance for app
func NewLogger() *LoggingService {
	return &LoggingService{
		logger: newLogger(),
	}
}

func newLogger() *zap.SugaredLogger {
	opts := zap.AddCallerSkip(1)
	logger, _ := zap.NewProduction(opts)
	return logger.Sugar()
}

func (l *LoggingService) Log(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

// Debug ...
func (l *LoggingService) Debug(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

// Info ...
func (l *LoggingService) Info(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

// Warn ...
func (l *LoggingService) Warn(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

// Error ...
func (l *LoggingService) Error(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}
