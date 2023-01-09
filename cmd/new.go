package cmd

import (
	"github.com/spf13/cobra"
	aftershock "github.com/thesubtlety/aftershock/internal"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Add terraform template files for a new provider or action",
	Long: `Quickly add a new provider or action
	Adds a folder and subfolder to the providers/ directory
	
	aftershock new <provider> <action-name>
	aftershock new github create-user`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Usage()
			return
		}
		aftershock.NewAction(args, conf)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
