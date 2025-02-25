package docker

import (
	"fmt"
	"os/exec"

	"github.com/AllanCapistrano/mysql-migrations/config"
)

var settings config.Settings

func init() {
	settings = config.GetSettings("mysql-migrations.json")
}

// Prepara a estrutura para a execução de comandos Data Definition Language (DDL)
func DdlCommand(query string) *exec.Cmd {
	passwordParameter := fmt.Sprintf("--password=%s", settings.DatabasePassword)

	return exec.Command(
		"docker", "exec", settings.ContainerName, "mysql", "-u", settings.DatabaseUser,
		passwordParameter, "-N", "-s", "-e", query,
	)
}

// Comando responsável por realizar o dump do banco de dados.
func DumpCommand(database string) *exec.Cmd {
	passwordParameter := fmt.Sprintf("--password=%s", settings.DatabasePassword)

	return exec.Command(
		"docker", "exec", settings.ContainerName, "mysqldump",
		"-u", settings.DatabaseUser, passwordParameter, database,
	)
}

// Comando responsável por restaurar o dump de um banco de dados.
func RestoreCommand(dump string, database string) *exec.Cmd {
	command := fmt.Sprintf(
		"cat %s | docker exec -i %s mysql -u %s --password=%s %s",
		dump,
		settings.ContainerName,
		settings.DatabaseUser,
		settings.DatabasePassword,
		database,
	)

	return exec.Command("sh", "-c", command)
}

// Comando responsável por realizar a migração a partir de um arquivo `.sql`.
func MigrateByFileCommand(filepath string, database string) *exec.Cmd {
	command := fmt.Sprintf(
		"docker exec -i %s mysql -u %s --password=%s %s < %s",
		settings.ContainerName,
		settings.DatabaseUser,
		settings.DatabasePassword,
		database,
		filepath,
	)

	return exec.Command("sh", "-c", command)
}

// Comando responsável por realizar a migração a partir de uma query.
func MigrateCommand(query string, database string) *exec.Cmd {
	command := DdlCommand(query)

	command.Args = append(command.Args, database)

	return command
}
