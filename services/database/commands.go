package database

import (
	"fmt"
	"os/exec"
)

// Prepara a estrutura para a execução de comandos Data Definition Language (DDL)
func ddlCommand(query string) *exec.Cmd {
	// TODO: Fazer com que algumas partes sejam configuradas, mas também ter um valor default
	return exec.Command(
		"docker", "exec", "conexa_mysql", "mysql", "-u", "root",
		"--password=root", "-N", "-s", "-e", query,
	)
}

// Comando responsável por realizar o dump do banco de dados.
func dumpCommand(database string) *exec.Cmd {
	// TODO: Fazer com que algumas partes sejam configuradas, mas também ter um valor default
	return exec.Command(
		"docker", "exec", "conexa_mysql", "mysqldump",
		"-u", "root", "--password=root", database,
	)
}

// Comando responsável por restaurar o dump de um banco de dados.
func restoreCommand(dump string, database string) *exec.Cmd {
	// TODO: Fazer com que algumas partes sejam configuradas, mas também ter um valor default
	command := fmt.Sprintf(
		"cat %s | docker exec -i conexa_mysql mysql -u root --password=root %s",
		dump,
		database,
	)

	return exec.Command("sh", "-c", command)
}

// Comando responsável por realizar a migração a partir de um arquivo `.sql`.
func migrateByFileCommand(filepath string, database string) *exec.Cmd {
    command := fmt.Sprintf(
        "docker exec -i conexa_mysql mysql -u root --password=root %s < %s",
        database,
        filepath,
    )

    return exec.Command("sh", "-c", command)
}