package services

import (
	"log"
	"os/exec"
	"strings"
)

const DATABASE_PREFIX = "opensev"

// Obtém todos os bancos de dados presentes no container do MySQL.
func getAllDatabases() string {
	// TODO: Fazer com que algumas partes sejam configuradas, mas também ter um valor default
	command := exec.Command(
		"docker", "exec", "conexa_mysql", "mysql", "-u", "root", 
		"--password=root", "-N", "-s", "-e", "SHOW DATABASES;",
	)

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
