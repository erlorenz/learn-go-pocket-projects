package main

import (
	"os"
	"time"

	"github.com/erlorenz/learn-go-pocket-projects/a-log-story/pocketlog"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelInfo, os.Stdout)

	lgr.Infof("A little copying is better than a little dependency.")
	lgr.Debugf("Make the (%d) zero value useful.", 0)
	lgr.Errorf("Errors are values. Documentation is for %s", "users")

	lgr.Infof("Hallo, %d %v", 2022, time.Now())
}
