package main

import (
	"log"
	"log/syslog"

	"github.com/TRedzepagic/compositelogger/logs"
	_ "github.com/go-sql-driver/mysql"
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

	systemlogger, _ := logs.NewSysLogger(syslog.LOG_NOTICE, log.LstdFlags)

	databaseLog := logs.NewDBLogger(logs.DatabaseConfiguration())
	defer databaseLog.Close()

	wantDebug := true

	// We can easily add another logger, for example:
	// loggerino := logs.NewFileLogger("newfileoutput")
	// All we need to do is to pass it as an argument to the NewCustomLogger function.

	log := logs.NewCustomLogger(wantDebug, filelogger1, filelogger2, stdoutLog, systemlogger, databaseLog)

	log.Info("info")
	log.Infof("%s", "infof")
	log.Warn("warn")
	log.Warnf("%s", "warnf")
	log.Debug("debug")
	log.Debugf("%s", "debugf")
	log.Error("error")
	log.Errorf("%s", "errorf")

}
