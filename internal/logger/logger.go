package logger

import (
	"log"
	"os"
)

var Log *log.Logger

func init() {
	f, err := os.OpenFile("/tmp/crr.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		// If file cannot be opened - skip logging
		Log = log.New(os.Stderr, "", 0)
		return
	}
	Log = log.New(f, "", log.LstdFlags)
}
