package pocketlog_test

import (
	"github.com/erlorenz/learn-go-pocket-projects/a-log-story/pocketlog"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("This is the debug %s", "level.")
	debugLogger.Infof("This is the info %s", "level.")
	debugLogger.Errorf("This is the error %s", "level.")
	// Output:
	// This is the debug level.
	// This is the info level.
	// This is the error level.
}

func ExampleLogger_Infof() {
	debugLogger := pocketlog.New(pocketlog.LevelInfo)
	debugLogger.Debugf("This is the debug %s", "level.")
	debugLogger.Infof("This is the info %s", "level.")
	debugLogger.Errorf("This is the error %s", "level.")
	// Output:
	// This is the info level.
	// This is the error level.
}

func ExampleLogger_Errorf() {
	debugLogger := pocketlog.New(pocketlog.LevelError)
	debugLogger.Debugf("This is the debug %s", "level.")
	debugLogger.Infof("This is the info %s", "level.")
	debugLogger.Errorf("This is the error %s", "level.")
	// Output:
	// This is the error level.
}
