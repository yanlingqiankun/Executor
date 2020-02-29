package logging

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

const (
	DebugLevel = logrus.DebugLevel
	InfoLevel = logrus.InfoLevel
	WarnLevel = logrus.WarnLevel
	ErrorLevel = logrus.ErrorLevel
)

type Fields logrus.Fields
type Logger struct {
	*logrus.Logger
}

func GetLogger(prefix string) *Logger {
	logger := logrus.New()
	logger.SetFormatter(&prefixed.TextFormatter{

	})
	logger.AddHook(&PrefixHook{prefix:prefix})
	return &Logger{
		Logger: logger,
	}
}

func GetLevel(levelString string) logrus.Level {
	switch levelString {
	case "debug": return DebugLevel
	case "info": return InfoLevel
	case "warn": return WarnLevel
	case "error": return ErrorLevel
	default:
		return InfoLevel
	}
}

type PrefixHook struct {
	prefix string
}

func (h *PrefixHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *PrefixHook) Fire(e *logrus.Entry) error {
	e.Data["prefix"] = h.prefix
	return nil
}

func (logger Logger) WithFields(fields Fields) *logrus.Entry {
	return logger.Logger.WithFields(logrus.Fields(fields))
}
