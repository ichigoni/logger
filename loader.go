package logger

import (
	"encoding/json"
	"log"
)

// Info is shortcut to Event logger
func Info(event, msg string) {
	Event("INFO", event, msg)
}

// Warn is shortcut to Event logger
func Warn(event, msg string) {
	Event("WARN", event, msg)
}

// Fatal is shortcut to Event logger
func Fatal(event, msg string) {
	Event("FATAL", event, msg)
}

// Event in json format
func Event(level, event, msg string) {
	thelog := New(level, event, msg)
	Println(thelog)
}

//Println is func actual print out log
func Println(thelog StructuredLog) {
	// todo: use sjson
	logJSON, err := json.Marshal(thelog)
	if err != nil {
		log.Println("Structured logger: Logger JSON Marshal failed !", err.Error())
	}

	if LogLevel == "" {
		LogLevel = "INFO"
	}
	if LogLevel == "INFO" {
		log.Println(LogLevel, string(logJSON))
		return
	}
	if LogLevel == "WARN" {
		if thelog.Level == "INFO" {
			return
		}
		// warn and fatal
		log.Println(LogLevel, string(logJSON))
		return
	}
	if LogLevel == thelog.Level {
		log.Println(LogLevel, string(logJSON))
	}
} 
