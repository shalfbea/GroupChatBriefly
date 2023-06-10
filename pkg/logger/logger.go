package logger

import (
	"os"

	"github.com/shalfbea/GroupChatBriefly/pkg/config"
	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
}

func InitLogger(config *config.Config) (logger Logger, err error) {
	logrus := logrus.New()
	logrus.Out = os.Stdout
	return logrus, nil
}
