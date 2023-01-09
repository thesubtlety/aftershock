package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	config "github.com/thesubtlety/aftershock/config"
	aftershock "github.com/thesubtlety/aftershock/internal"
)

var conf config.AftershockConfig

var rootCmd = &cobra.Command{
	Use:   "aftershock",
	Short: "A tool to make quick use of terraform providers",
	Long: `An tool to quickly leverage terraform providers and modules
Examples
	aftershock run github recon
	aftershock run splunk create-user`,
	//Run: func(cmd *cobra.Command, args []string) {},
}

func setupTF() {
	if conf.TfIgnorePath {
		conf.TfExecPath = aftershock.InstallTF()
		return
	}

	//Check PATH
	path, err := exec.LookPath("terraform")
	if err != nil {
		log.Println("terraform not found in path")
		// Previously installed with tfexec?
		tfpathb, err := ioutil.ReadFile(".tfexec")
		if err != nil {
			log.Println("unable to read .tfexec", err)
		}
		tfpath := string(tfpathb)
		tfpath, err = exec.LookPath(tfpath)
		if err != nil {
			// then force install
			log.Println("terraform not found at", tfpath)
			conf.TfExecPath = aftershock.InstallTF()
		}
		conf.TfExecPath = tfpath
		return
	}
	conf.TfExecPath = path
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(setupTF)
	rootCmd.PersistentFlags().BoolVar(&conf.TfIgnorePath, "ignore-path", false, "Ignore your installed version of terraform")
	rootCmd.PersistentFlags().BoolVar(&conf.Destroy, "destroy", false, "Destroy (undo) what the last plan you applied")
	rootCmd.PersistentFlags().BoolVar(&conf.Clean, "clean", false, "Delete provider terraform state files")
	rootCmd.PersistentFlags().BoolVar(&conf.Force, "force", false, "Don't prompt to delete state files")
	rootCmd.PersistentFlags().BoolVar(&conf.Verbose, "verbose", false, "Print standard terraform output")
	rootCmd.PersistentFlags().StringVar(&conf.DebugPath, "debug-path", "", "TF_LOG=DEBUG output file")
	rootCmd.PersistentFlags().StringVar(&conf.ProvidersPath, "providers-path", "providers", "Path to terraform templates folder")
}
