package logger

import (
	"fmt"
	"os"
	"time"
)

func (l *Logger) message(logType, format string, values ...interface{}) {
	if format[len(format)-1] != '\n' {
		format += "\n"
	}

	logFormat := ""
	if l.UseTimestamp {
		logFormat = fmt.Sprintf("%s %s: %s", time.Now().Format(time.RFC3339), logType, format)
	} else {
		logFormat = fmt.Sprintf("%s: %s", logType, format)
	}

	outputFd := os.Stdout
	if logType == LOG_FATAL || logType == LOG_WARNING {
		outputFd = os.Stderr
	}

	if l.TestFd != nil {
		outputFd = l.TestFd
	}

	fmt.Fprintf(outputFd, logFormat, values...)
}

func (l *Logger) Infof(format string, values ...interface{}) {
	l.message(LOG_INFO, format, values...)
}

func (l *Logger) Debugf(format string, values ...interface{}) {
	if !l.UseDebug {
		return
	}
	l.message(LOG_DEBUG, format, values...)
}

func (l *Logger) Warningf(format string, values ...interface{}) {
	l.message(LOG_WARNING, format, values...)
}

func (l *Logger) Fatalf(format string, values ...interface{}) {
	l.message(LOG_FATAL, format, values...)
	if l.TestFd != nil {
		return
	}
	os.Exit(1)
}
