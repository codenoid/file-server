package core

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type source struct {
	Type   string `yaml:"type"`
	Source string `yaml:"source"`
}

type config struct {
	IdentityType   string   `yaml:"identity_type"`
	IdentitySource string   `yaml:"identity_source"`
	Sources        []source `yaml:"sources"`
}

func loadConfig(configFile string) (config config) {

	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(b, &config)
	if err != nil {
		panic(err)
	}

	return
}
