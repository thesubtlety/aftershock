package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	config "github.com/thesubtlety/aftershock/config"
	aftershock "github.com/thesubtlety/aftershock/internal"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a provider action",
	Long: `Runs a terraform action module in the providers folder:
For example:
aftershock run github recon
aftershock run splunk create-user
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Usage()
			return
		}
		runCmd := checkArgs(args, conf)
		log.Printf("Running %s %s", runCmd.Provider, runCmd.Action)
		aftershock.ExecTF(runCmd, conf)
	},
}

func checkArgs(args []string, conf config.AftershockConfig) config.RunCmd {
	var runCmd config.RunCmd
	provider := args[0]
	action := args[1]
	action = strings.ReplaceAll(action, "-", "_")
	actionPath := filepath.Join(conf.ProvidersPath, provider, action)

	_, err := os.ReadDir(actionPath)
	if err != nil {
		log.Fatal("path not found: ", actionPath)
	}

	runCmd.Provider = provider
	runCmd.Action = action
	runCmd.Path = actionPath
	return runCmd
}

func init() {
	rootCmd.AddCommand(runCmd)
}
