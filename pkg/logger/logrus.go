package logger

import (
	"github.com/sirupsen/logrus"
)

type Logrus struct {
	logrus *logrus.Logger
}

func NewLogrus(isDebug bool) *Logrus {
	l := &Logrus{logrus.New()}

	if isDebug {
		// to enable debug
		l.logrus.Level = logrus.DebugLevel
	}
	//l.logrus.Formatter = &logrus.JSONFormatter{}
	return l
}

func (l Logrus) Debug(args ...interface{}) {
	l.logrus.Debugln(args...)
}

func (l Logrus) Info(args ...interface{}) {
	l.logrus.Infoln(args...)
}

func (l Logrus) Warn(args ...interface{}) {
	l.logrus.Warnln(args...)
}

func (l Logrus) Error(args ...interface{}) {
	l.logrus.Errorln(args...)
}

func (l Logrus) Fatal(args ...interface{}) {
	l.logrus.Fatalln(args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatalln(args...)
}

func Error(args ...interface{}) {
	logrus.Errorln(args...)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}
