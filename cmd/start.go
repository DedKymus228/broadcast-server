package cmd

import (
	"broadcast-server/internal/server"

	"github.com/spf13/cobra"
)

var port string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {

		server.Start(port)

	},
}

func init() {
	rootCmd.AddCommand(startCmd)

}
