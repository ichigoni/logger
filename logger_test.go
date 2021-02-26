package logger

import (
	// "bytes"
	"os"
	"os/exec"
	"regexp"
	"testing"
)

const (
	debugMsg    = "Debug msg"
	infoMsg     = "Info msg"
	warningMsg  = "Warning msg"
	errorMsg    = "Error msg"
	fatalMsg    = "Fatal msg"
)

var logMsg	= map[string]string{
	"DEBUG":    debugMsg,
	"INFO":     infoMsg,
	"WARN":		warningMsg,
	"ERROR":    errorMsg,
	"FATAL": 	fatalMsg,
}

func TestNew(t *testing.T) {
	levels := []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	regexps := []*regexp.Regexp{
		regexp.MustCompile(debugMsg),
		regexp.MustCompile(infoMsg),
		regexp.MustCompile(warningMsg),
		regexp.MustCompile(errorMsg),
		regexp.MustCompile(fatalMsg),
	}

	LogLevel = "DEBUG"
	for i, level := range levels {
		output, _ := logTemp(level)
		if !regexps[i].MatchString(output) {
			t.Errorf("The output doesn't contain the expected msg for the level: %s. [%s] %s", level, output, regexps[i])
		}
	}
}

func TestNew_unknownLevel(t *testing.T) {
	_, err := New("UNKNOWN", "", "pref")
	if err == nil {
		t.Error("The factory didn't return the expected error")
		return
	}
	if err != ErrorInvalidLogLevel {
		t.Errorf("The factory didn't return the expected error. Got: %s", err.Error())
	}
}

func TestNew_fatal(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		_, err := New("FATAL", "", "pref")
		if err != nil {
			t.Error("The factory returned an expected error:", err.Error())
			return
		}
		Fatal("crash!!!", "")
		os.Exit(1)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestNew_fatal")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

// logTemp is a temporary function to mock the result
func logTemp(level string) (string, error) {
	return Event(level, "pref", logMsg[level])
}