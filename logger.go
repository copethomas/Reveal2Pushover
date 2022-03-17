package main

import (
	"log"
	"os"
)

var (
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
)

func LoggingInit() {
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lmicroseconds)
	Warn = log.New(os.Stderr, "WARN: ", log.Ldate|log.Ltime|log.Lmicroseconds)
	Error = log.New(os.Stderr, "!!! - ERRO - !!! : ", log.Ldate|log.Ltime|log.Lmicroseconds)
}
