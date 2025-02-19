package logger

import (
	"fmt"
	"io"
	"log"

	general_i "github.com/beka-birhanu/vinom-interfaces/general"
)

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorReset  = "\033[0m"
)

type Logger struct {
	logger *log.Logger
}

func New(prefix, prefixColor string, writter io.Writer) (general_i.Logger, error) {
	logger := log.New(writter, fmt.Sprintf("%s[%s] %s", prefixColor, prefix, colorReset), log.LstdFlags|log.Lmsgprefix)
	return &Logger{
		logger: logger,
	}, nil
}

// Error implements i.Logger.
func (l *Logger) Error(message string) {
	l.logger.Printf("%s[ERROR]%s %s", colorRed, colorReset, message)
}

// Info implements i.Logger.
func (l *Logger) Info(message string) {
	l.logger.Printf("%s[INFO]%s %s", colorGreen, colorReset, message)
}

// Warning implements i.Logger.
func (l *Logger) Warning(message string) {
	l.logger.Printf("%s[WARNING]%s %s", colorYellow, colorReset, message)
}
