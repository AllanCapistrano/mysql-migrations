package database

import (
	"fmt"

	"github.com/AllanCapistrano/mysql-migrations/services/clog"
	"github.com/AllanCapistrano/mysql-migrations/services/docker"
)

// Lida com a remoção de um banco de dados
func handleDeleteDatabase(databaseName string) {
	fmt.Printf("Removendo o banco de dados '%s'", databaseName)

	query := fmt.Sprintf("DROP DATABASE %s;", databaseName)

	command := docker.DdlCommand(query)

	err := command.Run()
	if err != nil {
		message := fmt.Sprintf("Não foi possível remover o banco de dados %s - %v", databaseName, err)
		clog.Fatal(message, clog.ERROR)
	}
}

// Lida com a criação de um banco de dados vazio
func handleCreateDatabase(databaseName string) {
	fmt.Printf("Criando o banco de dados '%s'", databaseName)

	query := fmt.Sprintf("CREATE DATABASE %s;", databaseName)

	command := docker.DdlCommand(query)

	err := command.Run()
	if err != nil {
		message := fmt.Sprintf("Não foi possível criar o banco de dados %s - %v", databaseName, err)
		clog.Fatal(message, clog.ERROR)
	}
}

// Lida com a restauração de um banco de dados
func handleRestoreDatabase(databaseName string, snapshotFilePath string) {
	fmt.Printf("Restaurando o banco de dados '%s'", databaseName)

	command := docker.RestoreCommand(snapshotFilePath, databaseName)

	err := command.Run()
	if err != nil {
		message := fmt.Sprintf("Não foi possível restaurar o banco de dados %s - %v", databaseName, err)
		clog.Fatal(message, clog.ERROR)
	}
}

// Realiza o rollback de uma migração
func RollbackDatabase(databaseName string, snapshotFilePath string) {
	fmt.Printf("Iniciando processo de rollback do banco de dados '%s' utilizando a snapshot '%s'", databaseName, snapshotFilePath)

	handleDeleteDatabase(databaseName)
	handleCreateDatabase(databaseName)
	handleRestoreDatabase(databaseName, snapshotFilePath)

	fmt.Println("Processo de rollback finalizado!")
}
