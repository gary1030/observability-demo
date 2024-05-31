package log

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogFormatter() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}
