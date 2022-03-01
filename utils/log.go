package utils

import (
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	logFile, err := os.OpenFile("game_mail.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic("init log err")
	}
	Trace = log.New(logFile,"Trace:",log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(logFile,"Info:",log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(logFile,"Warning:",log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(logFile,"Error:",log.Ldate|log.Ltime|log.Lshortfile)

}
