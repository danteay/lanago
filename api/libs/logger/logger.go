package logger

import (
	"fmt"
	"log"
	"time"
)

// Logger is the structure that handle the log methods Debug,
// Info, Warning, Error and Log. This last Method is the general
// logging method that all other methods wrapps
type Logger struct{}

// logWriter is an structure to manage the wirter og the logging functions
type logWriter struct{}

// Writer is an override of the default logger writer to set the format nedded
// in the logs
func (w *logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(string(bytes))
}

// Init setup the logger configuration writer
func (l *Logger) Init() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
}

// Debug is the fuction that logs with a Debug tag, it wrapps the Log general
// function
func (l *Logger) Debug(c ...interface{}) {
	l.Log("Debug", c)
}

// Info is the fuction that logs with a Debug tag, it wrapps the Log general
// function
func (l *Logger) Info(c ...interface{}) {
	l.Log("Info", c)
}

// Warnin is the fuction that logs with a Debug tag, it wrapps the Log general
// function
func (l *Logger) Warnin(c ...interface{}) {
	l.Log("Warning", c)
}

// Error is the fuction that logs with a Debug tag, it wrapps the Log general
// function
func (l *Logger) Error(c ...interface{}) {
	l.Log("Error", c)
}

// Log is the general logging action, where 't' is a thex to identify the type
// of log that you are printing, and c is a collection of some other parameters
// of any type that will be printed as message of the log
func (l *Logger) Log(t string, c ...interface{}) {
	time := time.Now().Format("2006/01/02 - 15:04:05")

	msg := "[" + t + "] " + time

	log.Println(msg, c)
}
