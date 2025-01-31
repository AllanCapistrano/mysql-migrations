package services

import (
	"log"
	"strings"

	"github.com/AllanCapistrano/cnx-migrations/services/database"
)

const DATABASE_PREFIX = "opensev"

// Obtém todos os bancos de dados presentes no container do MySQL.
func getAllDatabases() string {
	command := database.DdlCommand("SHOW DATABASES;")

	output, err := command.Output()
	if err != nil {
		log.Fatal("Não foi possível obter todos os Banco de Dados")
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

// Obtém os bancos de dados `opensev_*`.
func GetDatabases() []string {
	databases := strings.Split(getAllDatabases(), "\n")

	return filterByPrefix(databases, DATABASE_PREFIX)
}
