package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Clear = &cobra.Command{
	Use:   "clear",
	Short: "Remove os arquivos de rollback",
	Long:  "Remove todos os arquivos de rollback que estão presentes no diretório '.rollback'",
	Run: func(cmd *cobra.Command, args []string) {
		clear()
	},
}

func clear() {
	fmt.Println("Removendo todos os arquivos de rollback...")

	// TODO: Implementar remoção de arquivos de rollback
}

func init() {
	Clear.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	Clear.Flags().BoolVarP(&customHelp, "help", "H", false, "Exibe as opções do comando 'clear'")
}
