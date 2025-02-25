package main

import (
	// "fmt"

	// "github.com/AllanCapistrano/mysql-migrations/services/database"
	"github.com/AllanCapistrano/mysql-migrations/cmd"
)

func main() {
	cmd.Execute()

	// databases := database.GetDatabases()

	// fmt.Println(databases)

	// database.DumpDatabase("opensev_recorrencia", ".")
	// database.ExecuteMigrationsByFile("opensev_recorrencia", "./migration_whatsapp.sql")
	// database.RollbackDatabase("opensev_recorrencia", "./snapshot_opensev_recorrencia_2025-02-04-1738714924138.sql")

	// query := `CREATE TABLE whatsapp_mensagens (
	// 	id INT AUTO_INCREMENT PRIMARY KEY,
	// 	telefone VARCHAR(15) NOT NULL,
	// 	mensagem TEXT NOT NULL,
	// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	// );`

	// database.ExecuteMigrationByQuery("opensev_recorrencia", query)

}
