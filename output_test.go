package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestOutput(t *testing.T) {
	logger := logrus.New()
	file, _ := os.OpenFile("application.log", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0666);
	logger.SetOutput(file)

	//logger.SetLevel(logrus.TraceLevel)
	logger.Debug("Ini info debug")
	logger.Info("Helo logic")
	logger.Warning("warning error")
}
