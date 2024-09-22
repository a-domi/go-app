package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	log_file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file=logFile err=%s", err.Error())
	}
	mutiLogFile := io.MultiWriter(os.Stdout, log_file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(mutiLogFile)
}
