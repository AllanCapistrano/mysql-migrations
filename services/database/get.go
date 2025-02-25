package database

import (
	"log"
	"strings"

	"github.com/AllanCapistrano/mysql-migrations/config"
	"github.com/AllanCapistrano/mysql-migrations/services"
	"github.com/AllanCapistrano/mysql-migrations/services/docker"
)

// Obtém todos os bancos de dados presentes no container do MySQL.
func getAllDatabases() string {
	command := docker.DdlCommand("SHOW DATABASES;")

	output, err := command.Output()
	if err != nil {
		log.Fatalf("Não foi possível obter todos os Banco de Dados - %v", err)
	}

	return string(output)
}

// Filtra os bancos de dados a partir de um prefixo.
func filterByPrefix(array []string, prefix string) []string {
	var result []string

	for _, str := range array {
		if strings.HasPrefix(str, prefix) {
			result = append(result, str)
		}
	}

	return result
}

// Obtém os bancos de dados, levando em consideração a whitelist e blacklist.
func GetDatabases() []string {
	databases := getAllValidDatabases()
	databasesInWhitelist := config.GetDatabasesInWhitelist()
	databasesInBlacklist := config.GetDatabasesInBlacklist()

	if len(databasesInWhitelist) > 0 {
		databases = databasesInWhitelist
	}

	if len(databasesInBlacklist) > 0 {
		databases = services.SliceDifference(databases, databasesInBlacklist)
	}

	return databases
}

// Obtém todos os bancos de dados válidos.
func getAllValidDatabases() []string {
	databases := strings.Split(getAllDatabases(), "\n")

	databasesPrefix := config.GetDatabasesPrefix()

	if databasesPrefix != "" {
		return filterByPrefix(databases, databasesPrefix)
	}

	return databases
}
