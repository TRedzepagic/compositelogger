package logs

import (
	"database/sql"
	"fmt"
	"log"
	"log/syslog"
	"os"
	"time"

	"github.com/mkmueller/golog"
)

//SuperLogger superclass
type SuperLogger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
	SetPrefix(s string)
}

//FileLogger subclass
type FileLogger struct {
	logger *golog.Logger
	fd     *os.File
}

//StdLogger subclass
type StdLogger struct {
	logger *golog.Logger
	fd     *os.File
}

//SysLogger subclass
type SysLogger struct {
	logger *log.Logger
	err    error
}

//DBLogger subclass
type DBLogger struct {
	database *sql.DB
	logger   *golog.Logger
}

//Close closes the IO for the logger
func (flogger *FileLogger) Close() {
	flogger.fd.Close()
}

//Println wraps the Println function of the logger
func (flogger FileLogger) Println(v ...interface{}) {
	flogger.logger.Println(v...)
}

//Printf wraps the Printf function of the logger
func (flogger FileLogger) Printf(format string, v ...interface{}) {
	flogger.logger.Printf(format, v...)
}

//SetPrefix wraps the SetPrefix function of the logger
func (flogger FileLogger) SetPrefix(s string) {
	flogger.logger.SetPrefix(s)
}

//NewFileLogger creates a new file logger (FileLogger)
func (flogger FileLogger) NewFileLogger(path string) FileLogger {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	flogger.logger = golog.New(path)
	flogger.fd = f
	return flogger
}

//NewStdLogger creates a new stdout logger (StdLogger)
func (stdlogger StdLogger) NewStdLogger() StdLogger {
	f, err := os.OpenFile("/dev/stdout", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	stdlogger.logger = golog.New(log.LstdFlags)
	stdlogger.fd = f
	return stdlogger
}

//Close closes the IO for the logger
func (stdlogger *StdLogger) Close() {
	stdlogger.fd.Close()
}

//Println wraps the Println function of the logger
func (stdlogger StdLogger) Println(v ...interface{}) {
	stdlogger.logger.Println(v...)
}

//Printf wraps the Printf function of the logger
func (stdlogger StdLogger) Printf(format string, v ...interface{}) {
	stdlogger.logger.Printf(format, v...)
}

//SetPrefix wraps the SetPrefix function of the logger
func (stdlogger StdLogger) SetPrefix(s string) {
	stdlogger.logger.SetPrefix(s)
}

//NewSysLogger creates a new system logger (SysLogger)
func (syslogger SysLogger) NewSysLogger(prio syslog.Priority, flag int) (*log.Logger, error) {
	syslogger.logger, syslogger.err = syslog.NewLogger(prio, flag)
	return syslogger.logger, syslogger.err
}

//SetPrefix wraps the SetPrefix function of the logger
func (syslogger SysLogger) SetPrefix(s string) {
	syslogger.logger.SetPrefix(s)
}

//Println wraps the Println function of the logger
func (syslogger SysLogger) Println(v ...interface{}) {
	syslogger.logger.Println(v...)
}

//Printf wraps the Printf function of the logger
func (syslogger SysLogger) Printf(format string, v ...interface{}) {
	syslogger.logger.Printf(format, v...)
}

//NewDBLogger routes the DB connection from "DatabaseConfiguration" to the DB logger, then creates it.
func (dblog DBLogger) NewDBLogger(db *sql.DB) DBLogger {
	dblog.database = db
	dblog.logger = golog.New()
	return dblog
}

//DatabaseConfiguration - sets up the DB connection, user:root, password:password (testing purposes)
func DatabaseConfiguration() *sql.DB {
	conn, err := sql.Open("mysql",
		"root:password@tcp(127.0.0.1:3306)/LOGGER")
	if err != nil {
		log.Print(err)
	}
	return conn
}

//ToDB writes to database
func (dblog DBLogger) ToDB(str string) {
	stmt, err := dblog.database.Prepare("INSERT INTO LOGS(PREFIX, DATE, TIME, TEXT) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Print(err)
	}
	prefix := dblog.logger.Prefix()
	date := fmt.Sprint(time.Now().Format("01-02-2006"))
	time := fmt.Sprint(time.Now().Format("15:04:05"))
	_, err = stmt.Exec(prefix, date, time, str)
	if err != nil {
		log.Print(err)
	}
}

//Println works differently for the database, converts the printed output to string, then passes it for database recording
func (dblog DBLogger) Println(v ...interface{}) {
	str := fmt.Sprint(v...)
	dblog.ToDB(str)
}

//Printf works differently for the database, converts the printed output to string, then passes it for database recording
func (dblog DBLogger) Printf(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	dblog.ToDB(str)
}

//SetPrefix wraps the SetPrefix function of the logger
func (dblog DBLogger) SetPrefix(s string) {
	dblog.logger.SetPrefix(s)
}

//CompositeLog structure was needed to implement the "true for debug, false for no debug" functionality.
type CompositeLog struct {
	slicelog []SuperLogger
	flag     bool
}

//NewCustomLogger adds all passed loggers into a slice of SuperLoggers (Variadic args...SuperLogger, proizvoljan broj loggera)
//We see here that we can add more loggers (we just need to pass them through)
func (biglog CompositeLog) NewCustomLogger(flag bool, args ...SuperLogger) CompositeLog {
	biglog.slicelog = args
	biglog.flag = flag
	return biglog
}

//The following functions change the prefix, then call each of the loggers' Print/f functions respectively.

//Info ..
func Info(composite CompositeLog, v ...interface{}) {
	for _, logger := range composite.slicelog {
		logger.SetPrefix("Info:")
		logger.Println(v)
	}

}

//Infof ..
func Infof(composite CompositeLog, format string, v ...interface{}) {

	for _, logger := range composite.slicelog {
		logger.SetPrefix("Infof:")
		logger.Printf(format, v)
	}

}

//Warn ..
func Warn(composite CompositeLog, v ...interface{}) {

	for _, logger := range composite.slicelog {
		logger.SetPrefix("WARNING:")
		logger.Println(v)
	}

}

//Warnf ..
func Warnf(composite CompositeLog, format string, v ...interface{}) {

	for _, logger := range composite.slicelog {
		logger.SetPrefix("WARNINGf:")
		logger.Printf(format, v)
	}

}

//Error ..
func Error(composite CompositeLog, v ...interface{}) {

	for _, logger := range composite.slicelog {
		logger.SetPrefix("Info:")
		logger.Println(v)
	}

}

//Errorf ..
func Errorf(composite CompositeLog, format string, v ...interface{}) {

	for _, logger := range composite.slicelog {
		logger.SetPrefix("ERROR:")
		logger.Printf(format, v)
	}

}

//Debug (skipped if flag == false)
func Debug(composite CompositeLog, v ...interface{}) {
	if composite.flag {
		for _, logger := range composite.slicelog {
			logger.SetPrefix("DEBUG:")
			logger.Println(v)
		}
	}

}

//Debugf (skipped if flag == false)
func Debugf(composite CompositeLog, format string, v ...interface{}) {
	if composite.flag {
		for _, logger := range composite.slicelog {
			logger.SetPrefix("ERROR:")
			logger.Printf(format, v)
		}
	}

}
