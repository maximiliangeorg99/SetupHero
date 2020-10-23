package environment

import (
	"path"

	"github.com/mitchellh/go-homedir"
)

// Environment struct
type environment struct {
	ProjectName    string
	ProjectVersion string
	ProjectUrl     string
	UserDirectory  string
	AppDirectory   string
	SettingsFile   string
}

// Environment singleton
var Environment *environment

// inits environment
func init() {
	userDir, err := homedir.Dir()
	if err != nil {
		panic("Unable to initialize environment, something went really wrong...")
	}

	appDir := path.Join(userDir, ".gordon")

	env := &environment{
		ProjectName:    "SetupHero",
		ProjectVersion: "0.0.1",
		ProjectUrl:     "http://github.com/maximiliangeorg99/SetupHero/",
		UserDirectory:  userDir,
		AppDirectory:   appDir,
		SettingsFile:   path.Join(appDir, "settings.yaml"),
	}
	Environment = env
}
