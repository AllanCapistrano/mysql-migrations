package cmd

import (
	"fmt"

	"github.com/AllanCapistrano/mysql-migrations/services"
	"github.com/AllanCapistrano/mysql-migrations/services/database"
	"github.com/spf13/cobra"
)

var Databases = &cobra.Command{
	Use:   "databases",
	Short: "Lista os bancos de dados que serão manipulados pelos comandos",
	Long:  "Lista todos os bancos de dados nos quais os comandos serão executados, leva em consideração os bancos de dados definidos na 'whitelist' e na 'blacklist'.\nOs bancos de dados definidos na 'whitelist' são os de maior precedência, ou seja, caso não esteja vazia, os comandos só serão executados nesses bancos de dados.\nOs bancos de dados definidos na 'blacklist' são filtrados, ou seja, caso não esteja vazia, os comando serão executados em todos os bancos de dados, menos naqueles que estão na 'blacklist'",
	Run: func(cmd *cobra.Command, args []string) {
		databases(args)
	},
}

func databases(args []string) {
	databases := database.GetDatabases()

	if len(chosenDatabases) > 0 {
		databases = chosenDatabases
	}

	if len(ignoredDatabases) > 0 {
		databases = services.SliceDifference(databases, ignoredDatabases)
	}

	database.HasRemainingDatabases(databases)

	fmt.Println(databases)
}

func init() {
	Databases.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	Databases.Flags().BoolVarP(&customHelp, "help", "H", false, "Exibe as opções do comando 'databases'")
}