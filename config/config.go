package config

type AftershockConfig struct {
	Clean        bool
	Force        bool
	Destroy      bool
	Verbose      bool
	DebugPath    string
	TfIgnorePath bool
	TfExecPath   string
}

type RunCmd struct {
	Provider string
	Action   string
	Path     string
}
