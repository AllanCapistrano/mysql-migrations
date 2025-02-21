package cmd

import (
	"fmt"
	"os"

	"github.com/AllanCapistrano/cnx-migrations/services"
	"github.com/AllanCapistrano/cnx-migrations/services/database"
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

	if len(databases) == 0 {
		fmt.Println("Não existem bancos de dados para realizar a migração")

		os.Exit(0)
	}

	fmt.Println(databases) // TODO: Remover

	// TODO: Realizar dump dos bancos de dados e exibir uma mensagem para cada banco de dados
}

func init() {
	Dump.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	Dump.Flags().BoolVarP(&customHelp, "help", "h", false, "Exibe as opções do comando 'dump'")
	Dump.Flags().StringSliceVarP(&chosenDatabases, "databases", "", []string{}, "Realiza o dump somente nos bancos de dados especificados. Para múltiplos bancos de dados, utilize vírgulas para separá-los.")
	Dump.Flags().StringArrayVarP(&chosenDatabases, "database", "D", []string{}, "Realiza o dump somente no banco de dado especificado. Para múltiplos bancos de dados, utilize a flag mais de uma vez.")
}
