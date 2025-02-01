package main

import (
	"fmt"

	"github.com/AllanCapistrano/cnx-migrations/services/database"
)

// "github.com/AllanCapistrano/cnx-migrations/cmd"

func main() {
	// cmd.Execute()

	databases := database.GetDatabases()

	fmt.Println(databases)

	// database.DumpDatabase("opensev_recorrencia", ".")
	database.RollbackDatabase("opensev_recorrencia", "./snapshot_opensev_recorrencia_2025-02-01-1738448077080.sql")
}
