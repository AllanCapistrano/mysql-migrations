package database

import "log"

// Realiza a migração a partir de um arquivo `.sql`
func ExecuteMigrationsByFile(databaseName string, filepath string) {
	log.Printf("Iniciando a migração do arquivo '%s' no banco de dados '%s'", filepath, databaseName)
	
	err := migrateByFileCommand(filepath, databaseName).Run()
	if err != nil {
		log.Fatalf("Não foi possível executar a migração do arquivo no banco de dados '%s' - %v", databaseName, err)
	}

	log.Println("Migração finalizada!")
}