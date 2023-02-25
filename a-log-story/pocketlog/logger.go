package pocketlog

import (
	"fmt"
	"io"
	"os"
)

type Logger struct {
	threshold Level
	output    io.Writer
}

// New returns you a logger, ready to log at the required threshold.
// Default output is stdout.
func New(threshold Level) *Logger {
	lgr := &Logger{
		threshold: threshold,
		output:    os.Stdout,
	}

	return lgr
}

// Debugf formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	// Default to stdout
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold > LevelDebug {
		return
	}

	l.logf(format, args...)
}

// Infof formats and prints a message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	// Default to stdout
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold > LevelInfo {
		return
	}

	l.logf(format, args...)
}

// Errorf formats and prints a message if the log level is error or higher.
func (l *Logger) Errorf(format string, args ...any) {
	// Default to stdout
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold > LevelError {
		return
	}

	l.logf(format, args...)
}

// logf prints the message to the output
func (l *Logger) logf(format string, args ...any) {

	_, _ = fmt.Printf(format+"\n", args...)
}
