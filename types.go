package logger

import (
	"errors"
	"os"
	"strings"
	"time"
)

// StructuredLog struct for structured loggin
type StructuredLog struct {
	Timestamp string `json:"timestamp,omitempty"`
	Service   string `json:"service,omitempty"`
	Thread    string `json:"thread,omitempty"`
	IP        string `json:"ip,omitempty"`
	Env       string `json:"env,omitempty"`
	Server    string `json:"server,omitempty"`

	Level   string `json:"level,omitempty"`
	Event   string `json:"event,omitempty"`
	Message string `json:"message,omitempty"`
}

const (
	// LevelDebug = 0
	LevelDebug = iota
	// LevelInfo = 1
	LevelInfo
	// LevelWarning = 2
	LevelWarning
	// LevelError = 3
	LevelError
	// LevelCritical = 4
	LevelCritical
)

var (
	//ErrorInvalidLogLevel is Error on invalid log level
	ErrorInvalidLogLevel = errors.New("invalid log level")

	//ErrorPrintln is Error on output error
	ErrorPrintln = errors.New("Logger print log failed")

	//ErrorJSONMarshal is Error on json marshal/decode
	ErrorJSONMarshal = errors.New("Logger JSON Marshal failed")

	logLevels = map[string]int{
		"DEBUG": LevelDebug,
		"INFO":  LevelInfo,
		"WARN":  LevelWarning,
		"ERROR": LevelError,
		"FATAL": LevelCritical,
	}

	//ServiceName is default service name
	ServiceName = ""

	//LogLevel is default log Level
	LogLevel = "INFO"

	// Env is default environment information
	Env = ""
)

// New is to log with a new StructuredLog struct
func New(level, event, msg string) (StructuredLog, error) {
	var thelog StructuredLog
	_, ok := logLevels[strings.ToUpper(level)]
	if !ok {
		return thelog, ErrorInvalidLogLevel
	}
	hostname, _ := os.Hostname()

	thelog.Timestamp = time.Now().Format(time.RFC3339)
	thelog.Server = hostname
	thelog.Env = Env
	thelog.Level = level
	thelog.Event = event
	thelog.Message = msg
	thelog.Service = ServiceName
	return thelog, nil
}
