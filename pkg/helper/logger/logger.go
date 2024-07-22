package logger

import (
	"log"
	"os"
)

type logger struct {
	l *log.Logger
}

type Logger interface {
	Error()
	Info()
	Warning()
}

var (
	myLogger *logger
)

func init() {
	errorFile, err := os.OpenFile("myLOG.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	myLogger.l = log.New(errorFile, "", log.Ldate|log.Ltime|log.Lshortfile)

}

func Error(err error) {
	myLogger.l.SetPrefix("ERROR: ")
	myLogger.l.Println(err)
}

func Info(info string) {
	myLogger.l.SetPrefix("INFO: ")
	myLogger.l.Println(info)
}

func Warning(warning string) {
	myLogger.l.SetPrefix("WARNING: ")
	myLogger.l.Println(warning)
}
