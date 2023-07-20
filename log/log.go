package log

import (
	syslog "log"

	"github.com/uhppoted/uhppoted-lib/log"
)

var queue = make(chan string, 8)

func Default() *syslog.Logger {
	return syslog.Default()
}

func Debugf(format string, args ...any) {
	log.Debugf(format, args...)
}

func Infof(format string, args ...any) {
	log.Infof(format, args...)
}

func Warnf(format string, args ...any) {
	log.Warnf(format, args...)
}

func Errorf(format string, args ...any) {
	log.Errorf(format, args...)
}

func Fatalf(format string, args ...any) {
	log.Fatalf(format, args...)
}
