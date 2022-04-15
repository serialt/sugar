package sugar

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v3"
)

// 读取配置文件filepath，使用config 接收
func LoadConfig(filepath string, Config *interface{}) (err error) {
	if filepath == "" {
		dir, _ := homedir.Dir()
		filepath = dir
	}
	filepath, err = homedir.Expand(filepath)
	if err != nil {
		fmt.Printf("Get config file failed: %v\n", err)
	}
	if !Exists(filepath) {
		fmt.Printf("File not exist, please check it: %v\n", filepath)
		os.Exit(8)
	}

	config, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Read config failed, please check the path: %v , err: %v\n", filepath, err)
	}
	err = yaml.Unmarshal(config, Config)
	if err != nil {
		fmt.Printf("Unmarshal to struct, err: %v", err)
	}
	// fmt.Printf("LoadConfig: %v\n", Config)
	fmt.Printf("Config path: %v\n", filepath)
	return
}

// 判断文件目录否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
