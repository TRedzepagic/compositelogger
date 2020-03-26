package logs

import (
	"log"
	"log/syslog"
)

// NewSysLogger creates a new system logger (SysLogger)
func NewSysLogger(prio syslog.Priority, flag int) (*log.Logger, error) {
	var syslogger SysLogger
	syslogger.logger, syslogger.err = syslog.NewLogger(prio, flag)
	return syslogger.logger, syslogger.err
}

// SetPrefix wraps the SetPrefix function of the logger
func (syslogger *SysLogger) SetPrefix(s string) {
	syslogger.logger.SetPrefix(s)
}

// Println wraps the Println function of the logger
func (syslogger *SysLogger) Println(v ...interface{}) {
	syslogger.logger.Println(v...)
}

// Printf wraps the Printf function of the logger
func (syslogger *SysLogger) Printf(format string, v ...interface{}) {
	syslogger.logger.Printf(format, v...)
}
