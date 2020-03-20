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

//FileLogger ..
type FileLogger struct {
	logger *golog.Logger
	fd     *os.File
}

//StdLogger ..
type StdLogger struct {
	logger *golog.Logger
	fd     *os.File
}

//SysLogger ..
type SysLogger struct {
	logger *log.Logger
	err    error
}

//DBLogger ..
type DBLogger struct {
	database *sql.DB
	logger   *golog.Logger
}

//DBData for easy parsing
type DBData struct {
	Prefix string
	Date   string
	Time   string
	Text   string
	Debug  string
}

//Close je wrapper za filelogger
func (flogger *FileLogger) Close() {
	flogger.fd.Close()
}

//Println je wrapper za filelogger
func (flogger FileLogger) Println(v ...interface{}) {
	flogger.logger.Println(v...)
}

//Printf je wrapper za filelogger
func (flogger FileLogger) Printf(format string, v ...interface{}) {
	flogger.logger.Printf(format, v...)
}

//SetPrefix je wrapper..
func (flogger FileLogger) SetPrefix(s string) {
	flogger.logger.SetPrefix(s)
}

//MakeLogger za FileLogger
func (flogger FileLogger) MakeLogger(path string) FileLogger {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	flogger.logger = golog.New(path)
	flogger.fd = f
	return flogger
}

//MakeLogger za StdLogger
func (stdlogger StdLogger) MakeLogger() StdLogger {
	f, err := os.OpenFile("/dev/stdout", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	stdlogger.logger = golog.New(log.LstdFlags)
	stdlogger.fd = f
	return stdlogger
}

//Close je wrapper za stdlogger
func (stdlogger *StdLogger) Close() {
	stdlogger.fd.Close()
}

//Println je wrapper za stdlogger
func (stdlogger StdLogger) Println(v ...interface{}) {
	stdlogger.logger.Println(v...)
}

//Printf je wrapper za StdLogger
func (stdlogger StdLogger) Printf(format string, v ...interface{}) {
	stdlogger.logger.Printf(format, v...)
}

//SetPrefix je wrapper za StdLogger
func (stdlogger StdLogger) SetPrefix(s string) {
	stdlogger.logger.SetPrefix(s)
}

//MakeLogger za SysLogger, pozivam funkciju iz sysloga
func (syslogger SysLogger) MakeLogger(prio syslog.Priority, flag int) (*log.Logger, error) {
	syslogger.logger, syslogger.err = syslog.NewLogger(prio, flag)
	return syslogger.logger, syslogger.err
}

//SetPrefix je wrapper za SysLogger
func (syslogger SysLogger) SetPrefix(s string) {
	syslogger.logger.SetPrefix(s)
}

//Println je wrapper za filelogger
func (syslogger SysLogger) Println(v ...interface{}) {
	syslogger.logger.Println(v...)
}

//Printf za SysLogger
func (syslogger SysLogger) Printf(format string, v ...interface{}) {
	syslogger.logger.Printf(format, v...)
}

//MakeLogger za DB ..
func (dblog DBLogger) MakeLogger(db *sql.DB) DBLogger {
	dblog.database = db
	dblog.logger = golog.New()
	return dblog
}

//DatabaseConfiguration je postavljanje konekcije na mySQL server
func DatabaseConfiguration() *sql.DB {
	conn, err := sql.Open("mysql",
		"root:password@tcp(127.0.0.1:3306)/LOGGER")
	if err != nil {
		log.Print(err)
	}
	return conn
}

//ToDB upisivanje u bazu
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

//Println za DB je wrapper za ToDB, da bi se zadovoljio i-face
func (dblog DBLogger) Println(v ...interface{}) {
	str := fmt.Sprint(v...)
	dblog.ToDB(str)
}

//Printf za DB je wrapper za ToDB, da bi se zadovoljio i-face
func (dblog DBLogger) Printf(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	dblog.ToDB(str)
}

//SetPrefix je wrapper za DBLogger
func (dblog DBLogger) SetPrefix(s string) {
	dblog.logger.SetPrefix(s)
}

//CompositeLog Potrebno radi debug flag..
type CompositeLog struct {
	slicelog []SuperLogger
	flag     bool
}

//NewCustomLogger ubacuje sve loggere u jedan logger (Variadic args...SuperLogger, proizvoljan broj loggera). Flag koristim da omogucim debug/f ili ne.
func (biglog CompositeLog) NewCustomLogger(flag bool, args ...SuperLogger) CompositeLog {
	biglog.slicelog = args
	biglog.flag = flag
	return biglog
}

//U funkcijama koje slijede mijenja se prefix, kako sam ja razumio da je poenta. Poziva se svacija SetPrefix i Print komanda ponaosob.

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

//Debug ..
func Debug(composite CompositeLog, v ...interface{}) {
	if composite.flag {
		for _, logger := range composite.slicelog {
			logger.SetPrefix("DEBUG:")
			logger.Println(v)
		}
	}

}

//Debugf ..
func Debugf(composite CompositeLog, format string, v ...interface{}) {
	if composite.flag {
		for _, logger := range composite.slicelog {
			logger.SetPrefix("ERROR:")
			logger.Printf(format, v)
		}
	}

}
