package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/AllanCapistrano/mysql-migrations/services/clog"
)

type Settings struct {
	DockerImageName  string   `json:"dockerImageName"`
	DatabaseUser     string   `json:"databaseUser"`
	DatabasePassword string   `json:"databasePassword"`
	Whitelist        []string `json:"whitelist"`
	Blacklist        []string `json:"blacklist"`
	DatabasesPrefix  string   `json:"databasesPrefix"`
}

// Obtém as configurações que foram definidas no arquivo de configurações. Caso
// não seja possível ler o arquivo, serão utilizadas as configurações padrão.
func GetSettings(fileName string) Settings {
	foundSettingsFile := true

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Não foi possível acessar o diretório home do usuário.", err)

		foundSettingsFile = false
	}

	filePath := filepath.Join(homeDir, ".config", "mysql-migrations", fileName)
	file, err := os.Open(filePath)
	if err != nil { // TODO: Colocar para os logs serem salvos em um arquivo
		message := fmt.Sprintf(
			"Não foi possível ler o arquivo '%s'! Serão utilizados os valores padrão.",
			fileName,
		)

		clog.Print(message, clog.WARNING)

		foundSettingsFile = false
	}

	defer file.Close()

	if foundSettingsFile {
		var settings Settings

		decoder := json.NewDecoder(file)
		err := decoder.Decode(&settings)
		if err != nil {
			message := fmt.Sprintf("Erro ao decodificar o JSON: %v", err)

			clog.Print(message, clog.ERROR)
		}

		return settings
	}

	return Settings{
		DockerImageName:  "mysql",
		DatabaseUser:     "root",
		DatabasePassword: "root",
		Whitelist:        []string{},
		Blacklist:        []string{},
		DatabasesPrefix:  "",
	}
}

// Retorna os bancos de dados que estão na whitelist.
func GetDatabasesInWhitelist() []string {
	settings := GetSettings("mysql-migrations.json")

	return settings.Whitelist
}

// Retorna os bancos de dados que estão na blacklist.
func GetDatabasesInBlacklist() []string {
	settings := GetSettings("mysql-migrations.json")

	return settings.Blacklist
}

// Retorna o prefixo dos bancos de dados.
func GetDatabasesPrefix() string {
	settings := GetSettings("mysql-migrations.json")

	return settings.DatabasesPrefix
}
