package logger

import "strings"

// Level defines all available log levels for log messages.
type Level int

// Log levels.
const (
	CRITICAL Level = iota
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
	PRINT
)

var levelNames = []string{
	"CRITICAL",
	"ERROR",
	"WARNING",
	"NOTICE",
	"INFO",
	"DEBUG",
	"PRINT",
}

// String returns the string representation of a logging level.
func (p Level) String() string {
	return levelNames[p]
}

// LogLevel returns the log level from a string representation.
func LogLevel(level string) Level {
	for i, name := range levelNames {
		if strings.EqualFold(name, level) {
			return Level(i)
		}
	}
	return ERROR
}

