package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zapcore.Field

var (
	Int = zap.Int
	String = zap.String
	Error = zap.Error
	Bool = zap.Bool
	Any = zap.Any
)

type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
}

type LoggerImpl struct {
	zap *zap.Logger
}

var customTimeFormat string

func New(level, namespace string) *LoggerImpl {
	if level == "" {
		level = LevelInfo
	}

	logger := LoggerImpl{
		zap: newZapLogger(level, time.RFC3339),
	}

	logger.zap = logger.zap.Named(namespace)

	zap.RedirectStdLog(logger.zap)

	return &logger
}


func (l *LoggerImpl) Debug(msg string, fields ...Field) {
	l.zap.Debug(msg, fields...)
}

func (l *LoggerImpl) Info(msg string, fields ...Field) {
	l.zap.Info(msg, fields...)
}

func (l *LoggerImpl) Warn(msg string, fields ...Field) {
	l.zap.Warn(msg, fields...)
}

func (l *LoggerImpl) Error(msg string, fields ...Field) {
	l.zap.Error(msg, fields...)
}

func (l *LoggerImpl) Fatal(msg string, fields ...Field) {
	l.zap.Fatal(msg, fields...)
}

func GetNamed(l Logger, name string) Logger {
	switch v := l.(type) {
	case *LoggerImpl:
		v.zap = v.zap.Named(name)
		return v
	default:
		l.Info("logger.GetNamed: invalid logger type")
		return l
	}
}

func WithFields(l Logger, fields ...Field) Logger {
	switch v := l.(type) {
	case *LoggerImpl:
		return &LoggerImpl{
			zap: v.zap.With(fields...),
		}
	default:
		l.Info("logger.WithFields: invalid logger type")
		return l
	}
}


func Cleanup(l Logger) error {
	switch v := l.(type) {
	case *LoggerImpl:
		return v.zap.Sync()
	default:
		l.Info("logger.Cleanup: invalid logger type")
		return nil
	}
}