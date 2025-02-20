package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Settings struct {
	ContainerName    string   `json:"containerName"`
	DatabaseUser     string   `json:"databaseUser"`
	DatabasePassword string   `json:"databasePassword"`
	Whitelist        []string `json:"whitelist"`
	Blacklist        []string `json:"blacklist"`
}

// Obtém as configurações que foram definidas no arquivo de configurações. Caso
// não seja possível ler o arquivo, serão utilizadas as configurações padrão.
func GetSettings(fileName string) Settings {
	foundSettingsFile := true

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("Não foi possível acessar o diretório home do usuário.", err)

		foundSettingsFile = false
	}

	filePath := filepath.Join(homeDir, ".config", "cnx-migrations", fileName)
	file, err := os.Open(filePath)
	if err != nil { // TODO: Colocar para os logs serem salvos em um arquivo
		// log.Printf(
		// 	"Não foi possível ler o arquivo '%s'! Serão utilizados os valores padrão. \n%v.\n",
		// 	fileName,
		// 	err,
		// )
		foundSettingsFile = false
	}

	defer file.Close()

	if foundSettingsFile {
		var settings Settings

		decoder := json.NewDecoder(file)
		err := decoder.Decode(&settings)
		if err != nil {
			log.Println("Erro ao decodificar o JSON:", err)
		}

		return settings
	}

	return Settings{
		ContainerName:    "conexa_mysql",
		DatabaseUser:     "root",
		DatabasePassword: "root",
		Whitelist:        []string{},
		Blacklist:        []string{},
	}
}

// Retorna os bancos de dados que estão na whitelist.
func GetDatabasesInWhitelist() []string {
	settings := GetSettings("cnx-migrations.json")

	return settings.Whitelist
}

// Retorna os bancos de dados que estão na blacklist.
func GetDatabasesInBlacklist() []string {
	settings := GetSettings("cnx-migrations.json")

	return settings.Blacklist
}
