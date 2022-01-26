package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

type SimpleHook struct {

}

func (s *SimpleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *SimpleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Simple hook", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SimpleHook{})
	logger.Warn("Hlo warn")
}
