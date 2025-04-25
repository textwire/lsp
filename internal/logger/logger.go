package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	Enabled bool

	// Level can be: error, warn, info, debug
	Level string

	// CustomPath for logs directory, not a file (optional)
	CustomPath string
}

var (
	Info   *log.Logger
	Error  *log.Logger
	Debug  *log.Logger
	config Config
)

func init() {
	config = Config{
		Enabled: true,
		Level:   "info",
	}

	initLoggers()
}
func New(filename string) *log.Logger {
	const fileMode = 0666

	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, fileMode)
	if err != nil {
		log.Panicf("The filepath %s is missing a file", filename)
	}

	return log.New(logfile, "[textwire lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}

func initLoggers() {
	if !config.Enabled {
		Info = log.New(io.Discard, "", 0)
		Error = log.New(io.Discard, "", 0)
		Debug = log.New(io.Discard, "", 0)
		return
	}

	logPath := getLogPath()
	if err := os.MkdirAll(filepath.Dir(logPath), 0755); err != nil {
		panic(fmt.Sprintf("Failed to create log directory: %v", err))
	}

	// Create log file with rotation
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file: %v", err))
	}

	// Set up loggers with appropriate prefixes
	Info = log.New(logFile, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(logFile, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(logFile, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func getLogsDirPath() string {
	if config.CustomPath != "" {
		return config.CustomPath
	}

	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), "textwire", "logs")
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), "Library", "Logs", "textwire")
	}

	return filepath.Join(os.Getenv("HOME"), ".local", "share", "textwire", "logs")
}

func getLogPath() string {
	baseDir := getLogsDirPath()
	return filepath.Join(baseDir, "textwire-lsp.log")
}
