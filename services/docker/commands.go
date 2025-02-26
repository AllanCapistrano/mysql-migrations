package docker

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/AllanCapistrano/mysql-migrations/config"
)

var settings config.Settings
var containerId string

func init() {
	settings = config.GetSettings("mysql-migrations.json")
	containerId = getContainerId(settings.DockerImageName)
}

// Obtém o ID do container a partir do nome da imagem Docker.
func getContainerId(imageName string) string {
	if imageName == "" {
		fmt.Println("É necessário informar o nome da imagem Docker, para que o container seja encontrado.")

		os.Exit(1)
	}

	command := fmt.Sprintf("docker ps --format '{{.ID}} {{.Image}}' | grep '%s' | awk '{print $1}' | head -n 1", imageName)

	output, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		fmt.Println("Erro ao buscar os containers")

		os.Exit(1)
	}

	containerId := string(output)

	if containerId == "" {
		fmt.Printf("Não foi encontrado nenhum container da imagem '%s'. Verifique se o nome da imagem está correto ou se o container está executando\n", imageName)

		os.Exit(0)
	}

	return strings.TrimSpace(containerId)
}

// Prepara a estrutura para a execução de comandos Data Definition Language (DDL)
func DdlCommand(query string) *exec.Cmd {
	passwordParameter := fmt.Sprintf("--password=%s", settings.DatabasePassword)

	return exec.Command(
		"docker", "exec", containerId, "mysql", "-u", settings.DatabaseUser,
		passwordParameter, "-N", "-s", "-e", query,
	)
}

// Comando responsável por realizar o dump do banco de dados.
func DumpCommand(database string) *exec.Cmd {
	passwordParameter := fmt.Sprintf("--password=%s", settings.DatabasePassword)

	return exec.Command(
		"docker", "exec", containerId, "mysqldump",
		"-u", settings.DatabaseUser, passwordParameter, database,
	)
}

// Comando responsável por restaurar o dump de um banco de dados.
func RestoreCommand(dump string, database string) *exec.Cmd {
	command := fmt.Sprintf(
		"cat %s | docker exec -i %s mysql -u %s --password=%s %s",
		dump,
		containerId,
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
		containerId,
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
