package cmd

import (
	"fmt"
	"os"

	"github.com/AllanCapistrano/cnx-migrations/services"
	"github.com/AllanCapistrano/cnx-migrations/services/database"
	"github.com/spf13/cobra"
)

var chosenDatabases []string
var ignoredDatabases []string

var Rollback = &cobra.Command{
	Use:   "rollback",
	Short: "Realiza o rollback da migração mais recente",
	Long:  "Realiza o rollback da migração mais recente a partir do último arquivo criado no diretório '.rollback'. Caso nenhum arquivo seja encontrado, o rollback não poderá ser realizado. Ao final do processo o arquivo de rollback é removido",
	Run: func(cmd *cobra.Command, args []string) {
		rollback()
	},
}

func rollback() {
	// TODO: Criar método para buscar os arquivos no diretório .rollback que fica em .config/cnx-migrations

	fileName := "nome_do_arquivo_de_rollback"

	fmt.Printf("Realizando rollback utilizando o arquivo '%s\n", fileName)

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

	// TODO: Apagar o arquivo utilizado para fazer o rollback, pois o mesmo já foi utilizado
}

func init() {
	Rollback.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	Rollback.Flags().BoolVarP(&customHelp, "help", "h", false, "Exibe as opções do comando 'rollback'")
	Rollback.Flags().StringSliceVarP(&chosenDatabases, "databases", "", []string{}, "Realiza o rollback somente nos bancos de dados especificados. Para múltiplos bancos de dados, utilize vírgulas para separá-los.")
	Rollback.Flags().StringArrayVarP(&chosenDatabases, "database", "D", []string{}, "Realiza o rollback somente no banco de dado especificado. Para múltiplos bancos de dados, utilize a flag mais de uma vez.")
	Rollback.Flags().StringSliceVarP(&ignoredDatabases, "no-databases", "", []string{}, "Realiza o rollback em todos bancos de dados, exceto nos especificados. Para múltiplos bancos de dados, utilize vírgulas para separá-los.")
	Rollback.Flags().StringArrayVarP(&ignoredDatabases, "no-database", "", []string{}, "Realiza o rollback em todos bancos de dados, exceto no especificado. Para múltiplos bancos de dados, utilize a flag mais de uma vez.")
}
