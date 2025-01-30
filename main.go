package main

import (
	"fmt"

	"github.com/AllanCapistrano/cnx-migrations/services"
)

// "github.com/AllanCapistrano/cnx-migrations/cmd"

func main() {
	// cmd.Execute()

	databases := services.GetDatabases()

	fmt.Println(databases)

	// services.DumpDatabase("opensev_recorrencia", ".")
	services.RollbackDatabase("opensev_recorrencia", "./snapshot_opensev_recorrencia_2025-01-30-1738278990153.sql")
}
