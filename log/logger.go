package log

import (
	"fmt"
	glog "log"
	"os"
)

type Logger struct {
	logger *glog.Logger
}

func New() *Logger {
	return &Logger{
		logger: glog.New(os.Stdout, "", glog.Ldate|glog.Ltime),
	}
}

func (l *Logger) Info(format string, args ...interface{}) {
	_format := fmt.Sprintf("[INFO] %s", format)
	l.logger.Printf(_format, args...)
}

func (l *Logger) Warn(format string, args ...interface{}) {
	_format := fmt.Sprintf("[WARN] %s", format)
	l.logger.Printf(_format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	_format := fmt.Sprintf("[ERROR] %s", format)
	l.logger.Printf(_format, args...)
}

func (l *Logger) Debug(format string, args ...interface{}) {
	_format := fmt.Sprintf("[DEBUG] %s", format)
	l.logger.Printf(_format, args...)
}

func (l *Logger) Fatal(format string, args ...interface{}) {
	_format := fmt.Sprintf("[Fatal] %s", format)
	l.logger.Fatalf(_format, args...)
}
