package logger

import (
	"os"
	"time"
)

// ServiceName name of service
var ServiceName = ""

// LogLevel Logging level, INFO, WARN, FATAL
var LogLevel = "INFO"

// Env Environment
var Env = ""

// StructuredLog struct for structured loggin
type StructuredLog struct {
	Timestamp string `json:"timestamp,omitempty"`
	Service   string `json:"service,omitempty"`
	Thread    string `json:"thread,omitempty"`
	IP        string `json:"ip,omitempty"`
	Env       string `json:"env,omitempty"`
	Server    string `json:"server,omitempty"`

	Level        string      `json:"level,omitempty"`
	Event        string      `json:"event,omitempty"`
	Message      string      `json:"message,omitempty"`
}

// New is to log with a new StructuredLog struct
func New(level, event, msg string) StructuredLog {
	var thelog StructuredLog
	hostname, _ := os.Hostname()

	thelog.Timestamp = time.Now().Format(time.RFC3339)
	thelog.Server = hostname
	thelog.Env = Env
	thelog.Level = level
	thelog.Event = event
	thelog.Message = msg
	thelog.Service = ServiceName
	return thelog
}
