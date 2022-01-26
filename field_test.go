package main

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithField("username","Vandy ahmad").Info("Helo world")
	logger.WithField("username","Vandy ahmad").WithField("name","vandy ").Info("Halo 2")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithFields(logrus.Fields{
		"username":"bambang",
		"name":"bambang pamungkas",
	}).Info("Halo cah gagah")
}
