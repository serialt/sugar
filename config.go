package sugar

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
)

type Service struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var service *Service

// func LoadConfig(filepath string) (err error) {
// 	if filepath == "" {
// 		return
// 	}
// 	filepath, _ = homedir.Expand(filepath)
// 	myconfig, err := ioutil.ReadFile(filepath)
// 	if err != nil {
// 		err = errors.New(fmt.Sprintf("Read config failed, please check the path: %v , err: %v\n", filepath, err))
// 		return
// 	}
// 	if err = yaml.Unmarshal(myconfig, &service); err != nil {
// 		err = errors.New(fmt.Sprintf("Unmarshal to struct, err: %v", err))
// 	}

// 	return
// }

// LoadConfig 读取配置文件filepath，使用out接收, 如果filepath为空，默认读取项目根目录config.yaml文件
func LoadConfig(filepath string) (byteConfg []byte, err error) {
	if filepath == "" {
		rootPath, _ := GetRootPath()
		filepath = fmt.Sprintf("%s/%s", rootPath, "config.yaml")

	} else {
		filepath, err = homedir.Expand(filepath)
		if err != nil {
			return
		}
		if !Exists(filepath) {
			err = errors.New(fmt.Sprintf("File not exist, please check it: %v\n", filepath))
			return
		}
	}

	byteConfg, err = ioutil.ReadFile(filepath)
	if err != nil {
		err = errors.New(fmt.Sprintf("Read config failed, please check the path: %v , err: %v\n", filepath, err))
	}
	return
}

// GetRootPath get the project root path
func GetRootPath() (rootPath string, err error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	rootPath = strings.Replace(dir, "\\", "/", -1)
	return
}

// Exists 判断文件目录否存在
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

func Env(key, def string) string {
	if x := os.Getenv(key); x != "" {
		return x
	}
	return def
}
