package log

import (
	"context"
	"github.com/Gssssssssy/ns-stored/pkg/config"
	"github.com/sirupsen/logrus"
	"os"
)

type Logger interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	WithFields(Fields) Logger
}

type LoggerWrapper struct {
	*logrus.Logger
}

type EntryWrapper struct {
	*logrus.Entry
}

func (w *EntryWrapper) WithFields(fields Fields) Logger {
	return &EntryWrapper{
		Entry: w.Entry.WithFields(logrus.Fields(fields)),
	}
}

func (w *LoggerWrapper) WithFields(fields Fields) Logger {
	return &EntryWrapper{
		Entry: w.Logger.WithFields(logrus.Fields(fields)),
	}
}

func NewLogger(cfg config.Provider) *LoggerWrapper {
	l := logrus.New()

	if cfg.GetBool("json_logs") {
		jsonFormatter := new(logrus.JSONFormatter)
		jsonFormatter.DisableTimestamp = false
		l.Formatter = jsonFormatter
	} else {
		txtFormatter := new(logrus.TextFormatter)
		txtFormatter.DisableColors = true
		txtFormatter.FullTimestamp = true
		l.Formatter = txtFormatter
	}
	l.Out = os.Stderr

	switch cfg.GetString("loglevel") {
	case "debug":
		l.Level = logrus.DebugLevel
	case "warning", "warn":
		l.Level = logrus.WarnLevel
	case "info":
		l.Level = logrus.InfoLevel
	case "error":
		l.Level = logrus.ErrorLevel
	default:
		l.Level = logrus.DebugLevel
	}
	return &LoggerWrapper{Logger: l}
}

var defaultLogger *LoggerWrapper

func init() {
	defaultLogger = NewLogger(config.Config())
}

type Fields map[string]interface{}

func (f Fields) With(k string, v interface{}) Fields {
	f[k] = v
	return f
}

func (f Fields) WithFields(f2 Fields) Fields {
	for k, v := range f2 {
		f[k] = v
	}
	return f
}

type Key int

var key Key

func FromContext(ctx context.Context) Logger {
	if ctx != nil {
		value, ok := ctx.Value(key).(Logger)
		if ok {
			return value
		}
	}
	return defaultLogger
}

func Infof(m context.Context, format string, args ...interface{}) {
	FromContext(m).Infof(format, args...)
}

func Errorf(m context.Context, format string, args ...interface{}) {
	FromContext(m).Errorf(format, args...)
}
