package logger

import (
	"encoding/json"
	"log"
)

// Debug is shortcut to Event logger
func Debug(event, msg string) (string, error) {
	return Event("DEBUG", event, msg)
}

// Info is shortcut to Event logger
func Info(event, msg string) (string, error) {
	return Event("INFO", event, msg)
}

// Warn is shortcut to Event logger
func Warn(event, msg string) (string, error)  {
	return Event("WARN", event, msg)
}

// Error is shortcut to Event logger
func Error(event, msg string) (string, error) {
	return Event("ERROR", event, msg)
}

// Fatal is shortcut to Event logger
func Fatal(event, msg string) (string, error) {
	return Event("FATAL", event, msg)
}

// Event in json format
func Event(level, event, msg string) (string, error) {
	thelog, _ := New(level, event, msg)
	return Println(thelog)
}

//Println is func actual print out log
func Println(thelog StructuredLog) (string, error) {
	// todo: use sjson
	logJSON, err := json.Marshal(thelog)
	if err != nil {
		log.Println("Structured logger: Logger JSON Marshal failed !", err.Error())
		return "", ErrorJSONMarshal
	}

	if LogLevel == "" {
		LogLevel = "INFO"
	}

	theLogLevel := logLevels[thelog.Level]
	logLevel := logLevels[LogLevel]

	if logLevel == LevelDebug {
		log.Println(LogLevel, string(logJSON))
		return string(logJSON), nil
	}

	if theLogLevel > logLevel {
		// warn and fatal
		log.Println(LogLevel, string(logJSON))
		return string(logJSON), nil
	}

	return "", ErrorPrintln
} 
