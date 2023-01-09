package aftershock

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/thesubtlety/aftershock/config"
)

func NewAction(args []string, conf config.AftershockConfig) {
	provider := args[0]
	action := args[1]
	action = strings.ReplaceAll(action, "-", "_")
	actionPath := filepath.Join(conf.ProvidersPath, provider, action)
	providerPath := filepath.Join(conf.ProvidersPath, provider)

	//if provider folder doesn't exist, create it and add template tf files
	_, err := os.Stat(providerPath)
	if errors.Is(err, os.ErrNotExist) {
		log.Println("Creating folder", providerPath)
		err := os.Mkdir(providerPath, 0764)
		if err != nil {
			log.Println(err)
		}
		addProviderTemplates(providerPath)
	}

	// if action folder doesn't exist, create it and add template tf file
	_, err = os.Stat(actionPath)
	if !errors.Is(err, os.ErrNotExist) {
		log.Fatalln("Path exists, exiting:", actionPath)
	} else {
		log.Println("Creating folder", actionPath)
		err := os.MkdirAll(actionPath, 0764)
		if err != nil {
			log.Println(err)
		}
		addActionTemplate(actionPath, action)
	}
}

func addActionTemplate(actionPath, action string) {
	actionTemplate := filepath.Join("tf_templates", "action.tf")
	actionDst := filepath.Join(actionPath, fmt.Sprintf("%s.tf", action))

	log.Println("Adding template to", actionDst)
	copyFile(actionTemplate, actionDst)
}

func addProviderTemplates(providerPath string) {
	tplFiles := []string{"provider.tf", "variables.tf", "terraform.tfvars", "terraform.tfvars.example", "README.md"}
	for _, f := range tplFiles {
		tplPath := filepath.Join("tf_templates", f)
		dstPath := filepath.Join(providerPath, f)
		log.Println("Adding template to", dstPath)
		copyFile(tplPath, dstPath)
	}
}
