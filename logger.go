package logger

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"io/ioutil"
	"log"
	"os"
)

const (
	debug = "DEBUG"
	info  = "INFO"
	warn  = "WARN"
	err = "ERROR"
)

type Logger struct {
	Out *log.Logger
	pkg string
}

func GetLogger(logFile string, pac string) *Logger {
	w, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file [%s]", logFile)
	}

	dc := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    1000,
		MaxBackups: 3,
		MaxAge:     28,
	}
	out := log.New(w, "", log.Ldate|log.Ltime)
	out.SetOutput(dc)

	return &Logger{
		Out: out,
		pkg: pac,
	}
}

func GetLoggerWithConfig(pac string, c *lumberjack.Logger) *Logger {
	w, err := os.OpenFile(c.Filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file [%s]", c.Filename)
	}

	out := log.New(w, "", log.Ldate|log.Ltime)
	out.SetOutput(c)

	return &Logger{
		Out: out,
		pkg: pac,
	}
}

func (l Logger) Debug(msg string) {
	l.Out.SetPrefix(fmt.Sprintf("[%s] [%s] ", debug, l.pkg))
	l.Out.Print(msg)
}

func (l Logger) Debugf(msg string, v ...interface{}) {
	l.Out.SetPrefix(fmt.Sprintf("[%s] [%s] ", debug, l.pkg))
	l.Out.Printf(msg, v...)
}

func (l Logger) Info(msg string) {
	l.Out.SetPrefix(fmt.Sprintf("[%s] [%s] ", info, l.pkg))
	l.Out.Print(msg)
}

func (l Logger) Infof(msg string, v ...interface{}) {
	l.Out.SetPrefix(fmt.Sprintf("[%s] [%s] ", info, l.pkg))
	l.Out.Printf(msg, v...)
}

func (l Logger) Warn(msg string) {
	l.Out.SetPrefix(fmt.Sprintf("[%s] [%s] ", warn, l.pkg))
	l.Out.Print(msg)
}

func (l Logger) Warnf(msg string, v ...interface{}) {
	l.Out.SetPrefix(fmt.Sprintf("[%s] [%s] ", warn, l.pkg))
	l.Out.Printf(msg, v...)
}

func (l Logger) Error(msg string) {
	l.Out.SetPrefix(fmt.Sprintf("[%s] [%s] ", err, l.pkg))
	l.Out.Print(msg)
}

func (l Logger) Errorf(msg string, v ...interface{}) {
	l.Out.SetPrefix(fmt.Sprintf("[%s] [%s] ", err, l.pkg))
	l.Out.Printf(msg, v...)
}

func (l Logger) Close() {
	l.Out.SetOutput(ioutil.Discard)
}
