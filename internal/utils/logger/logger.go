package logger

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/fatih/color"
)

// Logger provides colored logging for the CLI
type Logger struct {
	info *log.Logger
	err  *log.Logger
	slog *slog.Logger
}

// New creates a new instance of Logger with predefined colors and prefixes
func New() *Logger {
	catColor := color.New(color.FgCyan, color.Bold).SprintFunc()
	errColor := color.New(color.FgRed, color.Bold).SprintFunc()

	return &Logger{
		info: log.New(os.Stdout, "  "+catColor("🐱")+" ", 0),
		err:  log.New(os.Stderr, "  "+errColor("❌")+" ", log.Lshortfile),
		slog: slog.Default(),
	}
}

// Slog returns the underlying slog.Logger for infrastructure components
func (l *Logger) Slog() *slog.Logger {
	return l.slog
}

// Info logs an informational message in white
func (l *Logger) Info(format string, a ...any) {
	msgColor := color.New(color.FgWhite).SprintfFunc()
	if len(a) == 0 {
		l.info.Println(msgColor(format))
		return
	}
	l.info.Printf(msgColor(format), a...)
}

// Error logs an error message in red
func (l *Logger) Error(msg string, err error) {
	errText := color.New(color.FgRed, color.Bold).SprintFunc()

	if err != nil {
		l.err.Printf("%s: %v\n", errText(msg), err)
	} else {
		l.err.Printf("%s\n", errText(msg))
	}
}

// Raw prints the raw output (AI response) in a bright color to make it stand out
func (l *Logger) Raw(msg string) {
	artColor := color.New(color.FgHiYellow, color.Bold).SprintFunc()
	fmt.Println("\n" + artColor(msg) + "\n")
}

// Fatal logs a fatal error and exits the program
func (l *Logger) Fatal(msg string, err error) {
	l.Error(msg, err)
	os.Exit(1)
}
