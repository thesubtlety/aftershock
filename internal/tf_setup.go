package aftershock

import (
	"context"
	"io/ioutil"
	"log"
	"strings"

	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
)

func InstallTF() string {
	log.Println("Downloading terraform binary")
	installer := &releases.LatestVersion{
		Product: product.Terraform,
	}

	execPath, err := installer.Install(context.Background())
	if err != nil {
		log.Fatalf("error installing Terraform: %s", err)
	}

	err = ioutil.WriteFile(".tfexec", []byte(execPath), 0644)
	if err != nil {
		log.Println("unable to write .tfexec", err)
	}
	return execPath
}

func setupWorkspace(provider, modulePath, execPath string) *tfexec.Terraform {
	// Change to provider dir
	//providerPath := filepath.Join("providers", provider)
	tf, err := tfexec.NewTerraform(modulePath, execPath)
	if err != nil {
		log.Fatalf("error running NewTerraform: %s", err)
	}
	tf.SetAppendUserAgent("aftershock/v1")

	// Set up workspace
	workspaces, current, err := tf.WorkspaceList(context.Background())
	if err != nil {
		log.Fatalf("error running WorkspaceList: %s", err)
	}
	for _, workspace := range workspaces {
		if strings.Contains(workspace, provider) {
			err = tf.WorkspaceSelect(context.Background(), provider)
			if err != nil {
				log.Fatalf("error running WorkspaceSelect: %s", err)
			}
			current = workspace
			break
		}
	}
	if current != provider {
		err = tf.WorkspaceNew(context.Background(), provider)
		if err != nil {
			log.Fatalf("error running WorkspaceNew: %s", err)
		}
	}

	return tf
}
