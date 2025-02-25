package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var customHelp bool

var rootCmd = &cobra.Command{
	Use:   "mm",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bem vindo ao sistema de migrações do Conexa.")
		fmt.Printf("Versão: v1.0.0\n\n")
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func customHelpCommand() {
	rootCmd.InitDefaultHelpCmd()
	helpCmd := rootCmd.Commands()[2]
	helpCmd.Short = "Exibe informações sobre os comandos disponíveis"
}

func removeCompletionCommand() {
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
}

func init() {
	rootCmd.AddCommand(Migrate)
	rootCmd.AddCommand(Rollback)
	rootCmd.AddCommand(Dump)
	rootCmd.AddCommand(Clear)

	customHelpCommand()
	removeCompletionCommand()

	rootCmd.Flags().BoolVarP(&customHelp, "help", "h", false, "Utilize essa flag nos comandos para poder ver todas as opções disponíveis")
}
