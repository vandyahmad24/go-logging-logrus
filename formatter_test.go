package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestFormatter(t *testing.T) {
	logger := logrus.New()
	file, _ := os.OpenFile("application.log", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0666);
	logger.SetOutput(file)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.Info("Halo logger")
	logger.Warn("Halo logger")

}