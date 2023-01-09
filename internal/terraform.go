package aftershock

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/thesubtlety/aftershock/config"
)

func ExecTF(runCmd config.RunCmd, config config.AftershockConfig) {
	provider := runCmd.Provider
	modulePath := runCmd.Path

	if config.Clean {
		Clean(modulePath, config)
		return
	}

	moduleFiles := []string{"terraform.tfvars", "provider.tf", "variables.tf"}
	for _, f := range moduleFiles {
		from := filepath.Join(config.ProvidersPath, provider, f)
		to := filepath.Join(modulePath, f)
		//_, err := os.Stat(to)
		//if os.IsNotExist(err) {
		err := copyFile(from, to)
		if err != nil {
			log.Fatalln("error copying file", err)
		}
		//}
	}

	tfvars := tfexec.VarFile("terraform.tfvars")

	// Set up workspace for module
	tf := setupWorkspace(provider, modulePath, config.TfExecPath)

	// assign cli vars
	//https://github.com/porter-dev/switchboard/blob/82c227d9c378f3ec7d6bf739823882a92fd9f164/pkg/drivers/terraform/driver.go#L149

	if config.Verbose {
		tf.SetStdout(os.Stdout)
		tf.SetStderr(os.Stderr)
	}

	// init provider
	err := tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		log.Fatalf("error running Init: %s", err)
	}

	// set logging
	if config.DebugPath != "" {
		tf.SetLog("DEBUG")
		tf.SetLogPath(config.DebugPath)
	}

	// Plan tf
	_, err = tf.Plan(context.Background(), tfvars)
	if err != nil {
		log.Fatalf("error running Plan: %s", err)
	}

	if config.Destroy {
		log.Println("Destroying")
		err = tf.Destroy(context.Background(), tfvars)
		if err != nil {
			log.Fatalf("error running Destroy: %s", err)
		}
		return
	}

	log.Println("Applying")
	err = tf.Apply(context.Background(), tfvars)
	if err != nil {
		log.Fatalf("error running Apply: %s", err)
	}

	out, err := tf.Output(context.Background())
	if err != nil {
		log.Fatalf("error running Output: %s", err)
	}
	log.Println("Successfully applied plan")
	for k, v := range out {
		fmt.Printf("%+s\n", k)
		fmt.Printf("%s\n", v.Value)
	}

}

func Clean(modulePath string, conf config.AftershockConfig) {
	log.Println("Removing tf files - if you do this you can't auto clean up previous deploys with '--destroy'")

	if !conf.Force {
		fmt.Printf("Enter 'y' to proceed: ")
		var ok string
		fmt.Scanln(&ok)
		if ok != "y" {
			return
		}
	}

	stateFiles := []string{".terraform.tfstate.backup", ".terraform.lock.hcl", "terraform.tfvars", "provider.tf", "variables.tf"}
	stateDirs := []string{"terraform.tfstate.d/", ".terraform/"}

	// remove tf state files from provider directory
	for _, f := range stateFiles {
		r := filepath.Join(modulePath, f)
		os.Remove(r)
	}

	// remove state directory from provider directory
	for _, stateDir := range stateDirs {
		r := filepath.Join(modulePath, stateDir)
		os.RemoveAll(r)
	}

	log.Println("Finished cleaning up")
}

func copyFile(from, to string) error {
	input, err := os.ReadFile(from)
	if err != nil {
		return err
	}
	err = os.WriteFile(to, input, 0640)
	if err != nil {
		return err
	}
	return nil
}
