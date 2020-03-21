package main

import (
	"log/syslog"

	logs "github.com/TRedzepagic/compositelogger/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mkmueller/golog"
)

func main() {
	filepath1 := "logfile1"
	filepath2 := "logfile2"

	filelogger1 := logs.FileLogger.NewFileLogger(logs.FileLogger{}, filepath1)
	defer filelogger1.Close()

	filelogger2 := logs.FileLogger.NewFileLogger(logs.FileLogger{}, filepath2)
	defer filelogger2.Close()

	stdoutLog := logs.StdLogger.NewStdLogger(logs.StdLogger{})
	defer stdoutLog.Close()

	systemlogger, _ := logs.SysLogger.NewSysLogger(logs.SysLogger{}, syslog.LOG_NOTICE, golog.LstdFlags)

	databaseLog := logs.DBLogger.NewDBLogger(logs.DBLogger{}, logs.DatabaseConfiguration())

	wantDebug := true

	//We can easily add another logger, for example:
	//loggerino := golog.New("filelogtest") and all we need to do is to pass it as an argument to the NewCustomLogger function.

	log := logs.CompositeLog.NewCustomLogger(logs.CompositeLog{}, wantDebug, filelogger1, filelogger2, stdoutLog, systemlogger, databaseLog)
	logs.Info(log, "info")
	logs.Infof(log, "%s", "info")
	logs.Warn(log, "warn")
	logs.Warnf(log, "%s", "warnf")
	logs.Debug(log, "debug")
	logs.Debugf(log, "%s", "debugf")
	logs.Error(log, "error")
	logs.Errorf(log, "%s", "errorf")

}
