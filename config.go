package sugar

import (
	"io/ioutil"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v3"
)

// Load Read config from file.
func LoadConfig(configPath string, out interface{}) (err error) {
	configPath, err = homedir.Expand(configPath)
	if err != nil {
		return
	}
	byteConfig, err := ioutil.ReadFile(configPath)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(byteConfig, out)
	return
}
