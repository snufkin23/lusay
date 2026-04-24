package logger

import (
	"log"
	"os"
)

type Logger struct {
	info *log.Logger
	err  *log.Logger
}

func New() *Logger {
	return &Logger{
		info: log.New(os.Stdout, "🐱 [INFO]  ", log.LstdFlags),
		err:  log.New(os.Stderr, "❌ [ERROR] ", log.LstdFlags|log.Lshortfile),
	}
}

func (l *Logger) Info(format string, a ...any) {
	l.info.Printf(format, a...)
}

func (l *Logger) Error(msg string, err error) {
	if err != nil {
		l.err.Printf("%s | err=%v", msg, err)
		return
	}
	l.err.Printf("%s", msg)
}

func (l *Logger) Fatal(msg string, err error) {
	l.Error(msg, err)
	os.Exit(1)
}
