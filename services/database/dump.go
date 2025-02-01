package database

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Lida com a criação do nome do arquivo do backup do banco de dados.
func handleFileName(fileName string, path string) string {
	now := time.Now()

	year, month, day := now.Date()

	milliseconds := now.UnixMilli()

	currentDateWithHash := fmt.Sprintf("%d-%02d-%02d-%d", year, month, day, milliseconds)

	return fmt.Sprintf("%s/snapshot_%s_%s.sql", path, fileName, currentDateWithHash)
}

// Lida com a criação do arquivo que armazenará o backup do banco de dados.
func handleCreateOutputFile(fileName string, path string) (*os.File, error) {
	filePath := handleFileName(fileName, path)

	file, err := os.Create(filePath)

	if err != nil {
		return nil, fmt.Errorf("erro ao criar o arquivo de saída: %w", err)
	}

	return file, nil
}

// Realiza o dump de um banco de dados.
func DumpDatabase(databaseName string, outputPath string) {
	command := dumpCommand(databaseName)

	outputFile, err := handleCreateOutputFile(databaseName, outputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer outputFile.Close()

	command.Stdout = outputFile

	err = command.Run()
	if err != nil {
		log.Fatalf("Erro ao realizar o dump do banco de dados '%s' - %s", databaseName, err)
	}
}
