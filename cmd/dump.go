package cmd

import "github.com/spf13/cobra"

var Dump = &cobra.Command{
	Use:   "dump",
	Short: "Realiza o dump dos bancos de dados",
	Long:  "Realiza o dump de todos os bancos de dados (baseados na whitelist e blacklist, se estiverem preenchidas) organizando-os em diferentes diretórios",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

func init() {
	Dump.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	Dump.Flags().BoolVarP(&customHelp, "help", "h", false, "Exibe as opções do comando 'dump'")
}
