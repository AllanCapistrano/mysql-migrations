package services

import (
	"fmt"
	"log"
)

// Lida com a remoção de um banco de dados
func handleDeleteDatabase(database string) {
	query := fmt.Sprintf("DROP DATABASE %s;", database)

	command := handleDdlCommand(query)

	err := command.Run()
	if err != nil {
		log.Fatalf("Não foi possível remover o banco de dados %s - %s", database, err)
	}
}

// Lida com a criação de um banco de dados vazio
func handleCreateDatabase(database string) {
	query := fmt.Sprintf("CREATE DATABASE %s;", database)

	command := handleDdlCommand(query)

	err := command.Run()
	if err != nil {
		log.Fatalf("Não foi possível criar o banco de dados %s - %s", database, err)
	}
}

// Lida com a restauração de um banco de dados
func handleRestoreDatabase(database string, snapshotFilePath string) {
	command := handleRestoreCommand(snapshotFilePath, database)

	err := command.Run()
	if err != nil {
		log.Fatalf("Não foi possível restaurar o banco de dados %s - %s", database, err)
	}
}

// Realiza o rollback de uma migração
func RollbackDatabase(database string, snapshotFilePath string) {
	handleDeleteDatabase(database)
	handleCreateDatabase(database)
	handleRestoreDatabase(database, snapshotFilePath)
}
