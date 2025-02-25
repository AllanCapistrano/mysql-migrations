package database

import (
	"log"

	"github.com/AllanCapistrano/mysql-migrations/services/docker"
)

// Realiza a migração a partir de um arquivo `.sql`
func ExecuteMigrationsByFile(databaseName string, filepath string) {
	log.Printf("Iniciando a migração do arquivo '%s' no banco de dados '%s'", filepath, databaseName)

	err := docker.MigrateByFileCommand(filepath, databaseName).Run()
	if err != nil {
		log.Fatalf("Não foi possível executar a migração do arquivo no banco de dados '%s' - %v", databaseName, err)
	}

	log.Println("Migração finalizada!")
}

// Realiza a migração a partir de uma query
func ExecuteMigrationByQuery(databaseName string, query string) {
	log.Printf("Iniciando a migração no banco de dados '%s'", databaseName)

	err := docker.MigrateCommand(query, databaseName).Run()
	if err != nil {
		log.Fatalf("Não foi possível executar a migração do arquivo no banco de dados '%s' - %v", databaseName, err)
	}

	log.Println("Migração finalizada!")
}
