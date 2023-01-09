package cmd

import (
	"github.com/spf13/cobra"
	aftershock "github.com/thesubtlety/aftershock/internal"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available terraform providers and actions",
	Run: func(cmd *cobra.Command, args []string) {
		aftershock.ListActions(conf)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
