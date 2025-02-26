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
		fmt.Printf("MySQL Migrations - Versão v1.0.0\n\n")
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
	helpCmd := rootCmd.Commands()[3]
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
	rootCmd.AddCommand(Databases)

	customHelpCommand()
	removeCompletionCommand()

	rootCmd.Flags().BoolVarP(&customHelp, "help", "H", false, "Utilize essa flag nos comandos para poder ver todas as opções disponíveis")
}
