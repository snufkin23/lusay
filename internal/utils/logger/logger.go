package logger

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/fatih/color"
)

type Logger struct {
	info *log.Logger
	err  *log.Logger
	slog *slog.Logger
}

func New() *Logger {
	catColor := color.New(color.FgCyan, color.Bold).SprintFunc()
	errColor := color.New(color.FgRed, color.Bold).SprintFunc()

	return &Logger{
		info: log.New(os.Stdout, "  "+catColor("🐱")+" ", 0),
		err:  log.New(os.Stderr, "  "+errColor("❌")+" ", log.Lshortfile),
		slog: slog.Default(),
	}
}

func (l *Logger) Slog() *slog.Logger {
	return l.slog
}

func (l *Logger) Info(format string, a ...any) {
	msgColor := color.New(color.FgWhite).SprintfFunc()
	if len(a) == 0 {
		l.info.Println(msgColor(format))
		return
	}
	l.info.Printf(msgColor(format), a...)
}

func (l *Logger) Error(msg string, err error) {
	errText := color.New(color.FgRed, color.Bold).SprintFunc()

	if err != nil {
		l.err.Printf("%s: %v\n", errText(msg), err)
	} else {
		l.err.Printf("%s\n", errText(msg))
	}
}

func (l *Logger) Raw(msg string) {
	artColor := color.New(color.FgHiYellow, color.Bold).SprintFunc()
	fmt.Println("\n" + artColor(msg) + "\n")
}

func (l *Logger) Fatal(msg string, err error) {
	l.Error(msg, err)
	os.Exit(1)
}
