package logger

import (
	"context"
	"fmt"
	"log"
)

const (
	Info = iota
	Error
	Warning
	Debug
)

type Logger struct {
	ctx context.Context
}

func LogWithContext(ctx context.Context) *Logger {
	return &Logger{ctx: ctx}
}

func Log() *Logger {
	return &Logger{}
}

const loggingPattern = "[level: %d] [domain: %s] [message: %s] [detail: %v]"

func (l *Logger) Info(domain string, message string) {
	log.Println(fmt.Sprintf(loggingPattern, Info, domain, message, nil))
}

func (l *Logger) Error(domain string, message string, detail error) {
	log.Println(fmt.Sprintf(loggingPattern, Error, domain, message, detail))
}

