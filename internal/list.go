package aftershock

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/thesubtlety/aftershock/config"
)

type ProviderActions struct {
	Provider       string
	ProviderSource string
	Action         []string
	DataSources    []string
	Resources      []string
}

func ListActions(conf config.AftershockConfig) {
	providerPath := conf.ProvidersPath
	var list []string
	err := filepath.Walk(providerPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("error accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && strings.Count(path, string(os.PathSeparator)) > 1 {
			//smarter to chdir and `tfexec providers'
			list = append(list, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", providerPath, err)
		return
	}
	printList(list)
}

func printList(list []string) {
	actions := make(map[string][]string)
	for _, path := range list {
		p := strings.Split(path, string(os.PathSeparator))
		provider := p[1]
		action := p[2]
		actions[provider] = append(actions[provider], action)
	}
	for p, a := range actions {
		var provider ProviderActions
		provider.Provider = p
		provider.Action = a
		fmt.Printf("%s\n", provider.Provider)
		fmt.Printf("\t%s\n", strings.Join(provider.Action, "\n\t"))
	}
}
