package cmd

import (
	"fmt"
	"os"

	"github.com/AllanCapistrano/cnx-migrations/services"
	"github.com/spf13/cobra"
)

var sql string

var Migrate = &cobra.Command{
	Use:   "migrate",
	Short: "Realiza uma migração nos bancos de dados",
	Long:  "Realiza uma migração nos bancos de dados a partir de um arquivo '.sql' ou de uma query SQL. Durante o processo, são feitas cópias dos bancos de dados para caso seja necessário realizar o rollback, realizado através do comando 'cnx rollback'.",
	Run: func(cmd *cobra.Command, args []string) {
		migrate(args, sql)
	},
}

func migrate(args []string, sql string) {
	if sql != "" {
		migrateBySql(sql)
	}

	migrateByFile(args)
}

func migrateByFile(args []string) {
	if len(args) == 0 {
		fmt.Println("É necessário informar o nome do arquivo. Caso queira fazer a migração a partir de uma query SQL, utilize o parâmetro '--sql'")

		os.Exit(0)
	}

	if services.IsValidFile(args[0]) {
		fmt.Printf("Realizando a migração do arquivo '%s'\n", args[0]) // TODO: Colocar em um loop e informar o banco de dados.

		// TODO: Listar banco de dados.
		// TODO: Fazer o dump dos bancos de dados para possível rollback
		// TODO: Realizar a migração via arquivo

		os.Exit(0)
	}

	if services.CanBeSQLQuery(args[0]) {
		fmt.Println("Parece que você está tentando fazer uma migração a partir de uma query SQL, para isso, utilize o parâmetro '--sql'")

		os.Exit(0)
	}

	fmt.Println("Não foi possível prosseguir com a migração")

	os.Exit(1)
}

func migrateBySql(sql string) {
	if !services.IsSQLQuery(sql) {
		fmt.Println("A query informada não é válida, verifique e tente novamente")

		os.Exit(1)
	}

	fmt.Println("Realizando migração a partir da query SQL")

	// TODO: Listar banco de dados.
	// TODO: Fazer o dump dos bancos de dados para possível rollback
	// TODO: Realizar a migração via arquivo

	os.Exit(0)
}

func init() {
	Migrate.Flags().StringVarP(&sql, "sql", "S", "", "Query SQL para realizar a migração")
}
