package main

import (
	"log/syslog"

	logs "github.com/TRedzepagic/logger/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mkmueller/golog"
)

func main() {
	filepath1 := "logfile1"
	filepath2 := "logfile2"

	filelogger1 := logs.FileLogger.MakeLogger(logs.FileLogger{}, filepath1)
	defer filelogger1.Close()

	filelogger2 := logs.FileLogger.MakeLogger(logs.FileLogger{}, filepath2)
	defer filelogger2.Close()

	stdoutLog := logs.StdLogger.MakeLogger(logs.StdLogger{})
	defer stdoutLog.Close()

	systemlogger, _ := logs.SysLogger.MakeLogger(logs.SysLogger{}, syslog.LOG_NOTICE, golog.LstdFlags)

	databaseLog := logs.DBLogger.MakeLogger(logs.DBLogger{}, logs.DatabaseConfiguration())

	zelimDebug := true

	//Mogu se dodati proizvoljni loggeri.
	//NPR : loggerino := golog.New("filelogtest") bi pisao u navedeni file, a sve sto treba da uradimo jeste da ga dodamo u argumenta NewCustomLogger).

	log := logs.CompositeLog.NewCustomLogger(logs.CompositeLog{}, zelimDebug, filelogger1, filelogger2, stdoutLog, systemlogger, databaseLog)
	logs.Info(log, "OMAROVIC")
	logs.Infof(log, "%s", "OMAROVIC")
	logs.Warn(log, "warn")
	logs.Warnf(log, "%s", "warnf")
	logs.Debug(log, "debug")
	logs.Debugf(log, "%s", "debugf")
	logs.Error(log, "error")
	logs.Errorf(log, "%s", "errorf")

}
