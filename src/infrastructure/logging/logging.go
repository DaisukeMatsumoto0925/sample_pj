package logging

import (
	"fmt"
	"io"
	"log"
	"os"
)


type Logger struct {
	F *log.Logger
	I *log.Logger
	D *log.Logger
	E *log.Logger
}

type NullWriter struct{}

func (n NullWriter) Write(p []byte) (int, error) {
	return 0, nil
}

func NewLogger(env string) *Logger {
	var fatalWriter io.Writer = os.Stderr
	var infoWriter io.Writer = os.Stdout
	var errorWriter io.Writer = os.Stderr
	var debugWriter io.Writer = os.Stdout

	if env != "development" {
		debugWriter = NullWriter{}
	}

	return &Logger{
		F: log.New(fatalWriter, "[FATAL] ", log.Llongfile|log.Ldate|log.Lmicroseconds),
		E: log.New(errorWriter, "[ERROR] ", log.Llongfile|log.Ldate|log.Lmicroseconds),
		I: log.New(infoWriter, "[INFO]  ", log.Llongfile|log.Ldate|log.Lmicroseconds),
		D: log.New(debugWriter, "[DEBUG] ", log.Llongfile|log.Ldate|log.Lmicroseconds),
		}
}

func (l *Logger) Info(v ...interface{}) {
	l.I.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.I.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(v ...interface{}) {
	l.D.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.D.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.E.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.E.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.F.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.F.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}
