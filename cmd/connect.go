package cmd

import (
	"broadcast-server/internal/client"
	"fmt"

	"github.com/spf13/cobra"
)

var add string

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := client.Connect(add, port)
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.Flags().StringVarP(&add, "addres", "a", "localhost", "The address to connect to")
}
