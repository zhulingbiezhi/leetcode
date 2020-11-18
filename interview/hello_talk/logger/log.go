package logger

import (
	"fmt"
	"log"
	"time"
)

func Info(v ...interface{}) {
	log.Printf("[I] %s %+v \n", time.Now().Format("2006-01-02 15:04:05"), v)
}

func Infof(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	log.Printf("[I] %s %+v \n", time.Now().Format("2006-01-02 15:04:05"), s)
}

func Error(v ...interface{}) {
	log.Printf("[E] %s %+v \n", time.Now().Format("2006-01-02 15:04:05"), v)
}

func Errorf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	log.Printf("[E] %s %+v \n", time.Now().Format("2006-01-02 15:04:05"), s)
}
