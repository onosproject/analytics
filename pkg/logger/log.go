package logger

import (
	"fmt"
	"log"
	"os"
)

type LogLevel uint

const (
	ERROR LogLevel = iota
	WARN
	INFO
	DEBUG
)

var level LogLevel = ERROR

func Init(logFile string, logLevel LogLevel) {
	level = logLevel
	log.SetFlags(log.Ldate | log.Ltime)
	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			Warn("Unable to open %s for writing", logFile)
			Info("Will print log messages to stdout")
		} else {
			log.SetOutput(file)
		}
	}
}

func Error(message string, args ...interface{}) {
	newMessage := fmt.Sprintf(message, args...)
	log.Println("ERROR:", newMessage)
}
func Warn(message string, args ...interface{}) {
	newMessage := fmt.Sprintf(message, args...)
	log.Println("WARN:", newMessage)
}
func Info(message string, args ...interface{}) {
	newMessage := fmt.Sprintf(message, args...)
	log.Println("INFO:", newMessage)
}
func Debug(message string, args ...interface{}) {
	newMessage := fmt.Sprintf(message, args...)
	log.Println("DEBUG:", newMessage)
}
func Fatal(message string, args ...interface{}) {
	newMessage := fmt.Sprintf(message, args...)
	log.Fatalln("FATAL:", newMessage)
}
func Panic(message string, args ...interface{}) {
	newMessage := fmt.Sprintf(message, args...)
	log.Panicln("Panic:", newMessage)
}

func IfWarn() bool {
	return level >= WARN
}
func IfInfo() bool {
	return level >= INFO
}
func IfDebug() bool {
	return level == DEBUG
}
