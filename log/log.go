package log

import (
	"coroner/config"
	"github.com/sirupsen/logrus"
	"time"
)

func SetupLogger() {
	lvl, err := logrus.ParseLevel(config.Cfg.Logger.Level)
	if err != nil {
		lvl = logrus.ErrorLevel
	}
	logrus.SetLevel(lvl)

	if config.Cfg.Env == "dev" {
		logrus.SetReportCaller(true)
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339,
		})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	}
}
