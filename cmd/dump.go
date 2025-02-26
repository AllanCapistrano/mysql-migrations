package cmd

import (
	"fmt"

	"github.com/AllanCapistrano/mysql-migrations/services"
	"github.com/AllanCapistrano/mysql-migrations/services/database"
	"github.com/spf13/cobra"
)

var Dump = &cobra.Command{
	Use:   "dump",
	Short: "Realiza o dump dos bancos de dados",
	Long:  "Realiza o dump de todos os bancos de dados (baseados na whitelist e blacklist, se estiverem preenchidas) organizando-os em diferentes diretórios",
	Run: func(cmd *cobra.Command, args []string) {
		dump()
	},
}

func dump() {
	databases := database.GetDatabases()

	if len(chosenDatabases) > 0 {
		databases = chosenDatabases
	}

	if len(ignoredDatabases) > 0 {
		databases = services.SliceDifference(databases, ignoredDatabases)
	}

	database.HasRemainingDatabases(databases)

	fmt.Println(databases) // TODO: Remover

	// TODO: Realizar dump dos bancos de dados e exibir uma mensagem para cada banco de dados
}

func init() {
	Dump.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	Dump.Flags().BoolVarP(&customHelp, "help", "H", false, "Exibe as opções do comando 'dump'")
	Dump.Flags().StringSliceVarP(&chosenDatabases, "databases", "", []string{}, "Realiza o dump somente dos bancos de dados especificados. Para múltiplos bancos de dados, utilize vírgulas para separá-los.")
	Dump.Flags().StringArrayVarP(&chosenDatabases, "database", "D", []string{}, "Realiza o dump somente do banco de dado especificado. Para múltiplos bancos de dados, utilize a flag mais de uma vez.")
	Dump.Flags().StringSliceVarP(&ignoredDatabases, "no-databases", "", []string{}, "Realiza o dump de todos bancos de dados, exceto nos especificados. Para múltiplos bancos de dados, utilize vírgulas para separá-los.")
	Dump.Flags().StringArrayVarP(&ignoredDatabases, "no-database", "", []string{}, "Realiza o dump de todos bancos de dados, exceto no especificado. Para múltiplos bancos de dados, utilize a flag mais de uma vez.")
}
