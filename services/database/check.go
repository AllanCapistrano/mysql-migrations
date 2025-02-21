package database

import (
	"fmt"
	"os"
)

// Verifica se restou algum banco de dados para realizar a ação
func HasRemainingDatabases(databases []string) {
	if len(databases) == 0 {
		fmt.Println("Não existem bancos de dados para realizar a migração")

		os.Exit(0)
	}
}
