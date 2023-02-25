package pocketlog_test

import (
	"os"

	"github.com/erlorenz/learn-go-pocket-projects/a-log-story/pocketlog"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, os.Stdout)
	debugLogger.Debugf("Hello, %s", "world")
	// Output: Hello, world
}
