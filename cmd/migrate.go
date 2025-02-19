package cmd

import (
	"fmt"
	"os"

	"github.com/AllanCapistrano/cnx-migrations/services"
	"github.com/AllanCapistrano/cnx-migrations/services/database"
	"github.com/spf13/cobra"
)

var sql string
var chosenDatabases []string
var ignoredDatabases []string

var Migrate = &cobra.Command{
	Use:   "migrate",
	Short: "Realiza uma migração nos bancos de dados",
	Long:  "Realiza uma migração nos bancos de dados a partir de um arquivo '.sql' ou de uma query SQL. Durante o processo, são feitas cópias dos bancos de dados para caso seja necessário realizar o rollback, realizado através do comando 'cnx rollback'.",
	Run: func(cmd *cobra.Command, args []string) {
		migrate(args, sql)
	},
}

func migrate(args []string, sql string) {
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

	if sql != "" {
		migrateBySql(sql, databases)
	}

	migrateByFile(args, databases)
}

func migrateByFile(args []string, databases []string) {
	if len(args) == 0 {
		fmt.Println("É necessário informar o nome do arquivo. Caso queira fazer a migração a partir de uma query SQL, utilize o parâmetro '--sql'")

		os.Exit(0)
	}

	if !services.IsValidFile(args[0]) {
		if services.CanBeSQLQuery(args[0]) {
			fmt.Println("Parece que você está tentando fazer uma migração a partir de uma query SQL, para isso, utilize o parâmetro '--sql'")

			os.Exit(0)
		}

		fmt.Println("Não foi possível prosseguir com a migração")

		os.Exit(1)
	}

	fmt.Printf("Realizando a migração do arquivo '%s'\n", args[0]) // TODO: Colocar em um loop e informar o banco de dados.

	// TODO: Fazer o dump dos bancos de dados para possível rollback
	// TODO: Realizar a migração via arquivo

	os.Exit(0)
}

func migrateBySql(sql string, databases []string) {
	if !services.IsSQLQuery(sql) {
		fmt.Println("A query informada não é válida, verifique e tente novamente")

		os.Exit(1)
	}

	fmt.Println("Realizando migração a partir da query SQL") // TODO: Colocar em um loop e informar o banco de dados.

	// TODO: Fazer o dump dos bancos de dados para possível rollback
	// TODO: Realizar a migração via arquivo

	os.Exit(0)
}

func init() {
	Migrate.Flags().StringVarP(&sql, "sql", "S", "", "Especifica uma query SQL para realizar a migração. Mesmo que seja informado um arquivo, será realizada a migração da query SQL informada.")
	Migrate.Flags().StringSliceVarP(&chosenDatabases, "databases", "", []string{}, "Realiza a migração somente nos bancos de dados especificados. Para múltiplos bancos de dados, utilize vírgulas para separá-los. Ex: --databases db1,db2,db3")
	Migrate.Flags().StringArrayVarP(&chosenDatabases, "database", "D", []string{}, "Realiza a migração somente no banco de dado especificado. Para múltiplos bancos de dados, utilize a flag mais de uma vez. Ex: --database db1 --database db2")
	Migrate.Flags().StringSliceVarP(&ignoredDatabases, "no-databases", "", []string{}, "Realiza a migração em todos bancos de dados, exceto nos especificados. Para múltiplos bancos de dados, utilize vírgulas para separá-los. Ex: --no-databases db1,db2,db3")
	// TODO: Ver se é possível personalizar a mensagem da flag --help
}
