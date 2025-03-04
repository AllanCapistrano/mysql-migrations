package database

import (
	"fmt"

	"github.com/AllanCapistrano/mysql-migrations/services/clog"
	"github.com/AllanCapistrano/mysql-migrations/services/docker"
)

// Realiza a migração a partir de um arquivo `.sql`
func ExecuteMigrationsByFile(databaseName string, filepath string) {
	fmt.Printf("Iniciando a migração do arquivo '%s' no banco de dados '%s'", filepath, databaseName)

	err := docker.MigrateByFileCommand(filepath, databaseName).Run()
	if err != nil {
		message := fmt.Sprintf("Não foi possível executar a migração do arquivo no banco de dados '%s' - %v", databaseName, err)
		clog.Fatal(message, clog.ERROR)
	}

	fmt.Println("Migração finalizada!")
}

// Realiza a migração a partir de uma query
func ExecuteMigrationByQuery(databaseName string, query string) {
	fmt.Printf("Iniciando a migração no banco de dados '%s'", databaseName)

	err := docker.MigrateCommand(query, databaseName).Run()
	if err != nil {
		message := fmt.Sprintf("Não foi possível executar a migração do arquivo no banco de dados '%s' - %v", databaseName, err)
		clog.Fatal(message, clog.ERROR)
	}

	fmt.Println("Migração finalizada!")
}
