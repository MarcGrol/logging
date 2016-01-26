// +build appengine

package logging

import (
	"golang.org/x/net/context"

	"google.golang.org/appengine/log"
)

type (
	appengineLogger struct{}
)

func NewAppengineLogger() Logger {
	return &appengineLogger{}
}

func (logger *appengineLogger) Debug(c context.Context, format string, args ...interface{}) {
	log.Debugf(c, format, args...)
}

func (logger *appengineLogger) Info(c context.Context, format string, args ...interface{}) {
	log.Infof(c, format, args...)
}

func (logger *appengineLogger) Warning(c context.Context, format string, args ...interface{}) {
	log.Warningf(c, format, args...)
}

func (logger *appengineLogger) Error(c context.Context, format string, args ...interface{}) {
	log.Errorf(c, format, args...)
}

func (logger *appengineLogger) Critical(c context.Context, format string, args ...interface{}) {
	log.Criticalf(c, format, args...)
}

func init() {
	New = NewAppengineLogger
}
