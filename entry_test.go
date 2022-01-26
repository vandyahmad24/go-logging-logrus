package main

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	entry := logrus.NewEntry(logger)
	entry.WithField("username","vandy ahmad")
	entry.Info("Halo entry")
}
