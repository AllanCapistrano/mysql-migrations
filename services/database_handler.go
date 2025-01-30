package services

import "os/exec"

// Prepara a estrutura para a execução de comandos Data Definition Language (DDL)
func handleDdlCommand(query string) *exec.Cmd {
	// TODO: Fazer com que algumas partes sejam configuradas, mas também ter um valor default
	return exec.Command(
		"docker", "exec", "conexa_mysql", "mysql", "-u", "root",
		"--password=root", "-N", "-s", "-e", query,
	)
}

// Lida com a criação do comando responsável por realizar o dump do banco de dados.
func handleDumpCommand(database string) *exec.Cmd {
	// TODO: Fazer com que algumas partes sejam configuradas, mas também ter um valor default
	return exec.Command(
		"docker", "exec", "conexa_mysql", "mysqldump",
		"-u", "root", "--password=root", database,
	)
}
