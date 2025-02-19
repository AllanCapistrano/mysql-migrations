package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var customHelp bool

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
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(Migrate)

	customHelpCommand()

	rootCmd.Flags().BoolVarP(&customHelp, "help", "h", false, "Utilize essa flag nos comandos para poder ver todas as suas opções")
}

func customHelpCommand() {
	rootCmd.InitDefaultHelpCmd()
	helpCmd := rootCmd.Commands()[0] // O comando help é o primeiro comando adicionado
	helpCmd.Short = "Exibe informações sobre os comandos disponíveis"
}
