package clog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var filePath string

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
func Print(message string) {
	logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}

	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println(message)
}

// Escreve as mensagens de log em um arquivo customizado e para a execução do programa.
func Fatal(message string) {
	Print(message)
	os.Exit(1)
}
