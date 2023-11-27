package sugar

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v3"
)

// Load Read config from file.
func LoadConfig(configPath string, out interface{}) (err error) {
	configPath, err = homedir.Expand(configPath)
	if err != nil {
		return
	}
	// 当给定的配置文件不存在的时候，默认读取二进制文件所在目录下的config.yaml文件
	if !IsFile(configPath) {
		exePath, err := os.Executable()
		if err != nil {
			return err
		}
		configPath = filepath.Join(filepath.Dir(exePath), "config.yaml")
	}
	byteConfig, err := os.ReadFile(configPath)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(byteConfig, out)
	return
}

// IsFile
func IsFile(path string) bool {
	if path == "" {
		return false
	}

	if fi, err := os.Stat(path); err == nil {
		return !fi.IsDir()
	}
	return false
}
