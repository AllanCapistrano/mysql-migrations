package database

import (
	"log"
	"strings"
)

const DATABASE_PREFIX = "opensev"

// Obtém todos os bancos de dados presentes no container do MySQL.
func getAllDatabases() string {
	command := ddlCommand("SHOW DATABASES;")

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

// Obtém os bancos de dados `opensev_*`.
func GetDatabases() []string {
	databases := strings.Split(getAllDatabases(), "\n")

	return filterByPrefix(databases, DATABASE_PREFIX)
}
