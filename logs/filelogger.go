package logs

import (
	"fmt"
	"log"
	"os"
	"time"
)

// NewFileLogger creates a new file logger (FileLogger)
func NewFileLogger(path string) *FileLogger {
	var flogger FileLogger
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	flogger.fd = file
	return &flogger
}

// Println wraps the Println function of the logger
func (flogger *FileLogger) Println(v ...interface{}) {
	text := fmt.Sprintln(v...)
	date := fmt.Sprint(time.Now().Format("01-02-2006"))
	time := fmt.Sprint(time.Now().Format("15:04:05"))
	stringtowrite := flogger.fileloggerprefix + " " + date + " " + time + " " + text
	array := []byte(stringtowrite)
	flogger.fd.Write(array)
}

// Printf wraps the Printf function of the logger
func (flogger *FileLogger) Printf(format string, v ...interface{}) {
	text := fmt.Sprintf(format, v...)
	date := fmt.Sprint(time.Now().Format("01-02-2006"))
	time := fmt.Sprint(time.Now().Format("15:04:05"))
	stringtowrite := flogger.fileloggerprefix + " " + date + " " + time + " " + text + "\n"
	array := []byte(stringtowrite)
	flogger.fd.Write(array)

}

// SetPrefix wraps the SetPrefix function of the logger
func (flogger *FileLogger) SetPrefix(s string) {
	flogger.fileloggerprefix = s
}

// Close closes the IO for the logger
func (flogger *FileLogger) Close() {
	flogger.fd.Close()
}
