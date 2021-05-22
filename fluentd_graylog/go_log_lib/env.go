package log

import (
	"os"
	"strings"

	fluentd "github.com/joonix/log"
	"github.com/sirupsen/logrus"
)

func InitDefaultLoggerFromEnvironment(serviceName, version string) {

	logger := logrus.New()

	logFormatString := strings.ToLower(os.Getenv("LOG_FORMAT"))
	switch logFormatString {
	case "", "fluentd":
		logger.Formatter = &fluentd.FluentdFormatter{}
	case "pretty", "text":
		logger.Formatter = &logrus.TextFormatter{}
	default:
		logger.Formatter = &fluentd.FluentdFormatter{}
		logger.WithField("LOG_FORMAT", logFormatString).Error("Log Format invalid, using fluentd")
	}

	logger.Level = logrus.InfoLevel

	logLevelString := os.Getenv("LOG_LEVEL")
	if logLevelString != "" {
		level, err := logrus.ParseLevel(logLevelString)
		if err != nil {
			logger.WithField("LOG_LEVEL", logLevelString).Error("Log Level invalid, using debug")
			logger.Level = logrus.DebugLevel
		} else {
			logger.Level = level
		}
	}

	SetDefaultLog(logger.
		WithField("service", serviceName).
		WithField("version", version),
	)
}
