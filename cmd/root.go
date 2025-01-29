package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cnx",
	Short: "TODO",
	Long:  "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bem vindo ao sistema de migrações do Conexa")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal("Não foi possível executar o comando 'cnx'")
	}
}

func init() {
	rootCmd.AddCommand(Migrate)
}