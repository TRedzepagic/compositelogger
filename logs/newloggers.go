package logs

import (
	"log"
	"log/syslog"
	"os"
)

// NewSysLogger creates a new system logger (SysLogger)
func NewSysLogger(prio syslog.Priority, flag int) (*log.Logger, error) {
	var syslogger SysLogger
	syslogger.logger, syslogger.err = syslog.NewLogger(prio, flag)
	return syslogger.logger, syslogger.err
}

// NewFileLogger creates a new file logger (FileLogger)
func NewFileLogger(path string) FileLogger {
	var flogger FileLogger
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	flogger.logger = log.New(f, "Temporary_Prefix", log.LstdFlags)
	flogger.fd = f
	return flogger
}

// NewStdLogger creates a new stdout logger (StdLogger)
func NewStdLogger() StdLogger {
	var stdlogger StdLogger
	f, err := os.OpenFile("/dev/stdout", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	stdlogger.logger = log.New(f, "Temporary_Prefix", log.LstdFlags)
	stdlogger.fd = f
	return stdlogger
}
