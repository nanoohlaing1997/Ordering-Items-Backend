package log

import (
	"fmt"
	"io"
	"log/syslog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/nanoohlaing1997/online-ordering-items/env"
	log "github.com/sirupsen/logrus"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
	logrus_dailyrotate "github.com/yoannduc/logrus-dailyrotate"
	"gorm.io/gorm/logger"
)

var (
	logOnce         sync.Once
	logInstance     *log.Logger
	gormLogInstance logger.Interface
)

const (
	dailyChannel  = "daily"
	sysLogChannel = "syslog"
	microFormat   = "2006-01-02 15:04:05.000000000"
)

func GetFormatter(format string) log.Formatter {
	if format == "space" {
		return &nested.Formatter{
			HideKeys:        true,
			TimestampFormat: microFormat,
			NoColors:        false,
		}
	}
	return &log.JSONFormatter{}
}

// func dailyLog(logger *log.Logger, path string, format string) {
// 	hook, err := logrus_dailyrotate.New(
// 		path,
// 		30,
// 		GetFormatter(format),
// 		logger.Level,
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	logger.Hooks.Add(hook)
// }

func dailyLog(logger *log.Logger, relativePath string, format string) {
	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		panic(err)
	}

	hook, err := logrus_dailyrotate.New(
		absolutePath,
		30,
		GetFormatter(format),
		logger.Level,
	)
	if err != nil {
		panic(err)
	}

	logger.Hooks.Add(hook)
}

func sysLog(logger *log.Logger) {
	hook, err := logrus_syslog.NewSyslogHook("", "", syslog.Priority(logger.Level), "")
	if err != nil {
		logger.Warn("Unable to open syslog")
		logger.Out = os.Stdout
	} else {
		logger.AddHook(hook)
	}
}

func GetLogger() *log.Logger {
	logOnce.Do(func() {
		logInstance = log.New()
		env := env.GetEnviroment()
		logChannel := strings.ToLower(env.LogChannel)
		level, _ := log.ParseLevel(env.LogLevel)
		filePath := fmt.Sprintf("%s%s", env.LogFilePath, "ordering.log")

		logInstance.SetLevel(level)
		logInstance.SetOutput(io.Discard)
		logInstance.SetFormatter(GetFormatter(env.LogFormat))

		if logChannel == sysLogChannel {
			sysLog(logInstance)
		}

		if logChannel == dailyChannel {
			dailyLog(logInstance, filePath, env.LogFormat)
		}

		if filePath == "" {
			logInstance.SetOutput(os.Stdout)
		}
	})

	return logInstance
}

func dbLogLevel(level log.Level) logger.LogLevel {
	if level > 4 {
		return logger.Info
	}
	return logger.Silent
}

func GormLog() logger.Interface {
	env := env.GetEnviroment()
	log := GetLogger()

	gormLogInstance = logger.New(
		logInstance, logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  dbLogLevel(log.Level),
			IgnoreRecordNotFoundError: false,
			Colorful:                  !(env.LogChannel == "syslog"),
		})

	return gormLogInstance
}
