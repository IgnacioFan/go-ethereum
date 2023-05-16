package logger

import (
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	logger := log.New()

	// Set the desired log level
	logger.SetLevel(log.DebugLevel)

	// Create a file to write logs to
	logFile, err := os.OpenFile("tmp/logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	return logger
}
