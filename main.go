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

	filelogger1 := logs.NewFileLogger(filepath1)
	defer filelogger1.Close()

	filelogger2 := logs.NewFileLogger(filepath2)
	defer filelogger2.Close()

	stdoutLog := logs.NewStdLogger()
	defer stdoutLog.Close()

	systemlogger, _ := logs.NewSysLogger(syslog.LOG_NOTICE, golog.LstdFlags)

	databaseLog := logs.NewDBLogger(logs.DatabaseConfiguration())

	wantDebug := true

	//We can easily add another logger, for example:
	//loggerino := golog.New("filelogtest")
	//All we need to do is to pass it as an argument to the NewCustomLogger function.

	log := logs.NewCustomLogger(wantDebug, filelogger1, filelogger2, stdoutLog, systemlogger, databaseLog)
	logs.Info(log, "info")
	logs.Infof(log, "%s", "infof")
	logs.Warn(log, "warn")
	logs.Warnf(log, "%s", "warnf")
	logs.Debug(log, "debug")
	logs.Debugf(log, "%s", "debugf")
	logs.Error(log, "error")
	logs.Errorf(log, "%s", "errorf")

}
