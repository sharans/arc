package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
)

func Init() {
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stderr)
	if len(os.Getenv("LOG_LEVEL")) != 0 {
		level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
		if err == nil {
			logrus.SetLevel(level)
		}
	}
	logrus.SetLevel(logrus.InfoLevel)
}

func logger() *logrus.Logger {
	return logrus.StandardLogger()
}

func Debug(args ...interface{}) {
	sourced(logger()).Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	sourced(logger()).Debugf(format, args...)
}

func Info(args ...interface{}) {
	sourced(logger()).Info(args...)
}

func Infof(format string, args ...interface{}) {
	sourced(logger()).Infof(format, args...)
}

func Warn(args ...interface{}) {
	sourced(logger()).Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	sourced(logger()).Warnf(format, args...)
}

func Error(args ...interface{}) {
	sourced(logger(), true).Error(args...)
}

func Errorf(format string, args ...interface{}) {
	sourced(logger(), true).Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	sourced(logger(), true).Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	sourced(logger(), true).Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	sourced(logger(), true).Panic(args...)
}

func sourced(logger *logrus.Logger, includeStack ...bool) *logrus.Entry {
	var stack []string
	for i := 2; i <= 6; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok {
			stack = append(stack, formatStackEntry(file, line))
		}
	}

	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return logger.WithField("source", "<???>")
	}

	callerFunc := runtime.FuncForPC(pc).Name()

	sourcedLogger := logger.
		WithField("source", formatStackEntry(file, line)).
		WithField("methodName", callerFunc)
	if len(includeStack) > 0 && includeStack[0] {
		return sourcedLogger.WithField("stack", stack)
	}
	return sourcedLogger

}

func formatStackEntry(file string, line int) string {
	parts := strings.Split(file, "/")
	if len(parts) > 1 {
		file = strings.Join(parts[len(parts)-2:], "/")
	}

	return fmt.Sprintf("%s:%d", file, line)
}
