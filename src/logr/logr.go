package logr

import (
    "log"
    "os"
)

var logger *log.Logger

func init() {
    log.Println("logr.init")
    logFile, err := os.OpenFile("desr.log", os.O_CREATE|os.O_WRONLY, 0666)
    if (err != nil) {
	log.Fatalln("Failed to open log file: ", err)
    }
    logFile.Truncate(0)
    logger = log.New(logFile, "I ", log.Ldate|log.Ltime)
}

func LogI(args ...interface{}) {
    logger.SetPrefix("I ")
    logger.Println(args...)
}

func LogE(args ...interface{}) {
    logger.SetPrefix("E ")
    logger.Println(args...)
}
