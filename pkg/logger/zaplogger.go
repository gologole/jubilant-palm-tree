package mylogger

import (
	"fmt"
	"go.uber.org/zap"
	"runtime"
	"sync"
)

var GlobalLogger Logger

// Уровни логирования
const (
	debugLevel = iota
	infoLevel
	warnLevel
	errorLevel
)

var (
	mu    sync.Mutex
	level int = infoLevel
)

// Logger - структура логирования
type Logger struct {
	logger *zap.Logger
}

// NewLogger создает новый логгер.
func NewLogger() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	logger, _ := config.Build()
	GlobalLogger = Logger{logger: logger}
}

// SetLevel устанавливает уровень логирования.
func SetLevel(l int) {
	mu.Lock()
	defer mu.Unlock()
	level = l
}

// logMessage - логирует сообщение с уровнем логирования и указанием источника.
func (l *Logger) logMessage(level int, message string) {
	if level < int(l.logger.Level()) {
		return
	}

	_, file, line, ok := runtime.Caller(2)
	source := ""
	if ok {
		source = fmt.Sprintf("%s:%d", file, line)
	}

	switch level {
	case debugLevel:
		l.logger.Debug(message, zap.String("source", source))
	case infoLevel:
		l.logger.Info(message, zap.String("source", source))
	case warnLevel:
		l.logger.Warn(message, zap.String("source", source))
	case errorLevel:
		l.logger.Error(message, zap.String("source", source))
	}
}

// Debug - логирует сообщение на уровне DEBUG.
func (l *Logger) Debug(message string) {
	l.logMessage(debugLevel, message)
}

// Info - логирует сообщение на уровне INFO.
func (l *Logger) Info(message string) {
	l.logMessage(infoLevel, message)
}

// Warn - логирует сообщение на уровне WARN.
func (l *Logger) Warn(message string) {
	l.logMessage(warnLevel, message)
}

// Error - логирует сообщение на уровне ERROR.
func (l *Logger) Error(message string) {
	l.logMessage(errorLevel, message)
}
