package clog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var filePath string

type LogLevel int

const (
	INFO LogLevel = iota
	WARNING
	ERROR
)

func (logLevel LogLevel) String() string {
	switch logLevel {
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Não foi possível acessar o diretório home do usuário.", err)
	}

	currentDate := time.Now().Format("2006-01-02")
	fileName := fmt.Sprintf("%s.log", currentDate)

	filePath = filepath.Join(homeDir, ".config", "mysql-migrations", "logs", fileName)
}

// Escreve as mensagens de log em um arquivo customizado.
func Print(message string, level LogLevel) {
	logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}

	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.Ltime)
	log.Printf("[%s] %s\n", level, message)
}

// Escreve as mensagens de log em um arquivo customizado e para a execução do programa.
func Fatal(message string, level LogLevel) {
	Print(message, level)
	os.Exit(1)
}
