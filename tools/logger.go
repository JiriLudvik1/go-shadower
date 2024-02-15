package tools

import (
	"fmt"
	"os"
)

type Logger struct {
	file   *os.File
	config *LoggerConfig
}

type ILogger interface {
	Log(message string)
}

type LoggerConfig struct {
	FileName    string
	LogMatching bool
}

func NewLogger(config *LoggerConfig) (*Logger, error) {
	file, err := os.Create("log.txt")
	if err != nil {
		return nil, fmt.Errorf("failed to create log file: %w", err)
	}

	return &Logger{
		file:   file,
		config: config,
	}, nil
}

func (l *Logger) Log(message string) {
	fmt.Fprintln(l.file, message)
}
