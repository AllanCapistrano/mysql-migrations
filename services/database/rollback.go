package database

import (
	"fmt"
	"log"

	"github.com/AllanCapistrano/cnx-migrations/services/docker"
)

// Lida com a remoção de um banco de dados
func handleDeleteDatabase(databaseName string) {
	log.Printf("Removendo o banco de dados '%s'", databaseName)

	query := fmt.Sprintf("DROP DATABASE %s;", databaseName)

	command := docker.DdlCommand(query)

	err := command.Run()
	if err != nil {
		log.Fatalf("Não foi possível remover o banco de dados %s - %v", databaseName, err)
	}
}

// Lida com a criação de um banco de dados vazio
func handleCreateDatabase(databaseName string) {
	log.Printf("Criando o banco de dados '%s'", databaseName)

	query := fmt.Sprintf("CREATE DATABASE %s;", databaseName)

	command := docker.DdlCommand(query)

	err := command.Run()
	if err != nil {
		log.Fatalf("Não foi possível criar o banco de dados %s - %v", databaseName, err)
	}
}

// Lida com a restauração de um banco de dados
func handleRestoreDatabase(databaseName string, snapshotFilePath string) {
	log.Printf("Restaurando o banco de dados '%s'", databaseName)

	command := docker.RestoreCommand(snapshotFilePath, databaseName)

	err := command.Run()
	if err != nil {
		log.Fatalf("Não foi possível restaurar o banco de dados %s - %v", databaseName, err)
	}
}

// Realiza o rollback de uma migração
func RollbackDatabase(databaseName string, snapshotFilePath string) {
	log.Printf("Iniciando processo de rollback do banco de dados '%s' utilizando a snapshot '%s'", databaseName, snapshotFilePath)

	handleDeleteDatabase(databaseName)
	handleCreateDatabase(databaseName)
	handleRestoreDatabase(databaseName, snapshotFilePath)

	log.Println("Processo de rollback finalizado!")
}
