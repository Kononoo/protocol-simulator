package internal

import (
	"github.com/sirupsen/logrus"
	"os"
)

func InitLog(level string) {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	switch level {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	}

	logrus.SetOutput(os.Stdout)
}
