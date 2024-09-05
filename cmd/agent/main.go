package main

import (
	"log"
	"os"

	"golang.org/x/sys/windows/svc"

	"github.com/kdkumawat/mdm-poc/internal/agent"
)

const (
	logFilePath = "C:\\Windows\\Temp\\mdm-agent-service-poc.log"
)

func setupLogFile() *os.File {
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	return logFile
}

func main() {
	logFile := setupLogFile()
	defer logFile.Close()

	err := svc.Run("mdm-agent-service-poc", &agent.Service{})
	if err != nil {
		log.Println("failed to run service", err)
	}
}
