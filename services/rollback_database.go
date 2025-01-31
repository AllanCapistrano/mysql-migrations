package services

import (
	"fmt"
	"log"
)

// Lida com a remoção de um banco de dados
func handleDeleteDatabase(database string) {
	log.Printf("Removendo o banco de dados '%s'", database)

	query := fmt.Sprintf("DROP DATABASE %s;", database)

	command := ddlCommand(query)

	err := command.Run()
	if err != nil {
		log.Fatalf("Não foi possível remover o banco de dados %s - %s", database, err)
	}
}

// Lida com a criação de um banco de dados vazio
func handleCreateDatabase(database string) {
	log.Printf("Criando o banco de dados '%s'", database)

	query := fmt.Sprintf("CREATE DATABASE %s;", database)

	command := ddlCommand(query)

	err := command.Run()
	if err != nil {
		log.Fatalf("Não foi possível criar o banco de dados %s - %s", database, err)
	}
}

// Lida com a restauração de um banco de dados
func handleRestoreDatabase(database string, snapshotFilePath string) {
	log.Printf("Restaurando o banco de dados '%s'", database)

	command := restoreCommand(snapshotFilePath, database)

	err := command.Run()
	if err != nil {
		log.Fatalf("Não foi possível restaurar o banco de dados %s - %s", database, err)
	}
}

// Realiza o rollback de uma migração
func RollbackDatabase(database string, snapshotFilePath string) {
	log.Printf("Iniciando processo de rollback do banco de dados '%s' utilizando a snapshot '%s'", database, snapshotFilePath)

	handleDeleteDatabase(database)
	handleCreateDatabase(database)
	handleRestoreDatabase(database, snapshotFilePath)

	log.Println("Processo de rollback finalizado!")
}
