package logs

import (
	"database/sql"
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
	fileloggerprefix string
	fd               *os.File
}

// StdLogger subclass
type StdLogger struct {
	stdloggerprefix string
	fd              *os.File
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
