package logs

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// NewDBLogger routes the DB connection from "DatabaseConfiguration" to the DB logger, then creates it.
func NewDBLogger(db *sql.DB) *DBLogger {
	var dblog DBLogger
	dblog.database = db
	return &dblog
}

// DatabaseConfiguration - sets up the DB connection
func DatabaseConfiguration() *sql.DB {
	conn, err := sql.Open("mysql", "compositelogger:Mystrongpassword1234$@tcp(127.0.0.1:3306)/LOGGER") //Configuration described in README.md
	if err != nil {
		log.Print(err)
	}
	return conn
}

// ToDB writes to database
func (dblog *DBLogger) ToDB(str string) {
	stmt, err := dblog.database.Prepare("INSERT INTO LOGS(PREFIX, DATE, TIME, TEXT) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Print(err)
	}
	date := fmt.Sprint(time.Now().Format("01-02-2006"))
	time := fmt.Sprint(time.Now().Format("15:04:05"))
	_, err = stmt.Exec(dblog.prefix, date, time, str)
	if err != nil {
		log.Print(err)
	}
}

// Println (for DB) converts the printed output to string, then passes it for database recording
func (dblog *DBLogger) Println(v ...interface{}) {
	str := fmt.Sprint(v...)
	dblog.ToDB(str)
}

// Printf (for DB) converts the printed output to string, then passes it for database recording
func (dblog *DBLogger) Printf(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	dblog.ToDB(str)
}

// SetPrefix for DB
func (dblog *DBLogger) SetPrefix(s string) {
	dblog.prefix = s
}

// Close closes the DB connection
func (dblog *DBLogger) Close() {
	dblog.database.Close()
}
