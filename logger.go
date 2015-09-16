package logger

// Import packages
import (
	"fmt"
	//"io"
	"log"
	"log/syslog"
)

type LoggerConfig struct {
	Application string
	Output      WriterL
}

const (
	DEBUG    = "DEBUG"
	INFO     = "INFO"
	NOTICE   = "NOTICE"
	WARNING  = "WARNING"
	ERROR    = "ERROR"
	CRITICAL = "CRITICAL"
)

// Info class, Contains all the info on what has to logged, time is the current time, Module is the specific module
// For which we are logging, level is the state, importance and type of message logged,
// Message contains the string to be logged, format is the format of string to be passed to sprintf
type Info struct {
	Time        string
	Application string
	Level       string
	Message     string
	format      string
}

// Logger class that is an interface to user to log messages, Module is the module for which we are testing
// worker is variable of Worker class that is used in bottom layers to log the message
type Logger struct {
	Application string
	logger      WriterL
}

type WriterL interface {
	Alert(m string) (err error)
	Close() error
	Crit(m string) (err error)
	Debug(m string) (err error)
	Emerg(m string) (err error)
	Err(m string) (err error)
	Info(m string) (err error)
	Notice(m string) (err error)
	Warning(m string) (err error)
	Write(b []byte) (int, error)
}

// Returns a proper string to be outputted for a particular info
func (r *Info) Output() string {
	msg := fmt.Sprintf(r.format, r.Time, r.Level, r.Message)
	return msg
}

func New(config LoggerConfig) *Logger {
	//logWriter, err := syslog.Dial("", "localhost", syslog.LOG_DEBUG, config.Application)
	if config.Output == nil {
		writer, err := syslog.New(syslog.LOG_DEBUG, config.Application)

		if err != nil {
			log.Fatal("error")
		}
		config.Output = writer
	}
	return &Logger{Application: config.Application, logger: config.Output}
}

func (l *Logger) Close() {
	l.logger.Close()
}

// Critical logs a message at a Critical Level
func (l *Logger) Critical(message string) {
	l.logger.Crit(message)
}

// Error logs a message at Error level
func (l *Logger) Error(message string) {
	l.logger.Err(message)
}

// Warning logs a message at Warning level
func (l *Logger) Warning(message string) {
	l.logger.Warning(message)
}

// Notice logs a message at Notice level
func (l *Logger) Notice(message string) {
	l.logger.Notice(message)
}

// Info logs a message at Info level
func (l *Logger) Info(message string) {
	l.logger.Info(message)
}

// Debug logs a message at Debug level
func (l *Logger) Debug(message string) {
	l.logger.Debug(message)
}
