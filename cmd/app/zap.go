package app

import (
	"github.com/ThreeDotsLabs/watermill"
	"go.uber.org/zap"
)

// ZapLoggerAdapter implements the watermill.LoggerAdapter interface
type ZapLoggerAdapter struct {
	logger *zap.Logger
}

// NewZapLoggerAdapter creates a new ZapLoggerAdapter
func NewZapLoggerAdapter(logger *zap.Logger) *ZapLoggerAdapter {
	return &ZapLoggerAdapter{logger: logger}
}

// Debug logs a debug message
func (z *ZapLoggerAdapter) Debug(msg string, fields watermill.LogFields) {
	z.logger.Debug(msg, zap.Any("fields", fields))
}

// Info logs an info message
func (z *ZapLoggerAdapter) Info(msg string, fields watermill.LogFields) {
	z.logger.Info(msg, zap.Any("fields", fields))
}

// Error logs an error message
func (z *ZapLoggerAdapter) Error(msg string, err error, fields watermill.LogFields) {
	z.logger.Error(msg, zap.Error(err), zap.Any("fields", fields))
}

// Trace logs a trace message
func (z *ZapLoggerAdapter) Trace(msg string, fields watermill.LogFields) {
	z.logger.Debug("[TRACE] "+msg, zap.Any("fields", fields))
}

// With creates a new logger with additional fields
func (z *ZapLoggerAdapter) With(fields watermill.LogFields) watermill.LoggerAdapter {
	newLogger := z.logger.With(zap.Any("fields", fields))
	return &ZapLoggerAdapter{logger: newLogger}
}
