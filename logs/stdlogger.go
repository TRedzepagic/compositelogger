package logs

import (
	"fmt"
	"log"
	"os"
	"time"
)

// NewStdLogger creates a new stdout logger (StdLogger) by assigning the opened /dev/stdout to the file descriptor.
func NewStdLogger() *StdLogger {
	var stdlogger StdLogger
	file, err := os.OpenFile("/dev/stdout", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	stdlogger.fd = file
	return &stdlogger
}

// Println directly uses the file descriptors' Write()
func (stdlogger *StdLogger) Println(v ...interface{}) {
	text := fmt.Sprintln(v...)
	date := fmt.Sprint(time.Now().Format("01-02-2006"))
	time := fmt.Sprint(time.Now().Format("15:04:05"))
	stringtowrite := stdlogger.stdloggerprefix + " " + date + " " + time + " " + text
	array := []byte(stringtowrite)
	stdlogger.fd.Write(array)
}

// Printf directly uses the file descriptors' Write()
func (stdlogger *StdLogger) Printf(format string, v ...interface{}) {
	text := fmt.Sprintf(format, v...)
	date := fmt.Sprint(time.Now().Format("01-02-2006"))
	time := fmt.Sprint(time.Now().Format("15:04:05"))
	stringtowrite := stdlogger.stdloggerprefix + " " + date + " " + time + " " + text + "\n"
	array := []byte(stringtowrite)
	stdlogger.fd.Write(array)
}

// SetPrefix sets the prefix of the logger
func (stdlogger *StdLogger) SetPrefix(s string) {
	stdlogger.stdloggerprefix = s
}

// Close closes the IO for the logger
func (stdlogger *StdLogger) Close() {
	stdlogger.fd.Close()
}
