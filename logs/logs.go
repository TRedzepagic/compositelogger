package logs

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// SuperLogger superclass
type SuperLogger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
	SetPrefix(s string)
}

// FileLogger subclass
type FileLogger struct {
	logger *log.Logger
	fd     *os.File
}

// StdLogger subclass
type StdLogger struct {
	logger *log.Logger
	fd     *os.File
}

// SysLogger subclass
type SysLogger struct {
	logger *log.Logger
	err    error
}

// DBLogger subclass
type DBLogger struct {
	database *sql.DB
	prefix   string
}

// Close closes the IO for the logger
func (flogger *FileLogger) Close() {
	flogger.fd.Close()
}

// Println wraps the Println function of the logger
func (flogger FileLogger) Println(v ...interface{}) {
	flogger.logger.Println(v...)
}

// Printf wraps the Printf function of the logger
func (flogger FileLogger) Printf(format string, v ...interface{}) {
	flogger.logger.Printf(format, v...)
}

// SetPrefix wraps the SetPrefix function of the logger
func (flogger FileLogger) SetPrefix(s string) {
	flogger.logger.SetPrefix(s)
}

// Close closes the IO for the logger
func (stdlogger *StdLogger) Close() {
	stdlogger.fd.Close()
}

// Println wraps the Println function of the logger
func (stdlogger StdLogger) Println(v ...interface{}) {
	stdlogger.logger.Println(v...)
}

// Printf wraps the Printf function of the logger
func (stdlogger StdLogger) Printf(format string, v ...interface{}) {
	stdlogger.logger.Printf(format, v...)
}

// SetPrefix wraps the SetPrefix function of the logger
func (stdlogger StdLogger) SetPrefix(s string) {
	stdlogger.logger.SetPrefix(s)
}

// SetPrefix wraps the SetPrefix function of the logger
func (syslogger SysLogger) SetPrefix(s string) {
	syslogger.logger.SetPrefix(s)
}

// Println wraps the Println function of the logger
func (syslogger SysLogger) Println(v ...interface{}) {
	syslogger.logger.Println(v...)
}

// Printf wraps the Printf function of the logger
func (syslogger SysLogger) Printf(format string, v ...interface{}) {
	syslogger.logger.Printf(format, v...)
}

// Println works differently for the database, converts the printed output to string, then passes it for database recording
func (dblog DBLogger) Println(v ...interface{}) {
	str := fmt.Sprint(v...)
	dblog.ToDB(str)
}

// Printf works differently for the database, converts the printed output to string, then passes it for database recording
func (dblog DBLogger) Printf(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	dblog.ToDB(str)
}

// SetPrefix for DB
func (dblog *DBLogger) SetPrefix(s string) {
	dblog.prefix = s
}
