// Package container ...
//
// File:  log
//
// Description: log
//
// Date: 2021/5/24 4:31 下午
package container

import (
	"lib/src/github.com/sirupsen/logrus"
	"os"
)

func InitLog() {
	rootPath, _ := GetRootPath()

	logConfig, err := GetLogConfig(rootPath + "config/log")

	if err != nil {
		panic("init log failed")
	}

	logConfig.API = getLogInstance(logrus.DebugLevel)
	logConfig.Business = getLogInstance(logrus.DebugLevel)
	logConfig.Exception = getLogInstance(logrus.DebugLevel)
	logConfig.Redis = getLogInstance(logrus.DebugLevel)
	logConfig.SQL = getLogInstance(logrus.DebugLevel)
}

// getLogInstance ...
func getLogInstance(level logrus.Level, fileName string) (*logrus.Logger, error) {

	if level <= 0 {
		level = logrus.DebugLevel
	}

	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, err
	}

	logger := logrus.New()

	logger.Out = src

	logger.SetLevel(level)

	logger.SetFormatter(&logrus.TextFormatter{})

	return logger, nil
}
