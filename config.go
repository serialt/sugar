package sugar

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// 读取配置文件filepath，使用config 接收
func LoadConfig(filepath string, config interface{}) (err error) {

	myconfig, err := ioutil.ReadFile(filepath)
	if err != nil {
		logSugar.Infof("Read config failed, please check the path: %v , err: %v", filepath, err)
		return
	}
	err = yaml.Unmarshal(myconfig, &config)
	if err != nil {
		logSugar.Infof("Unmarshal to struct failed: %v", err)
		return
	}
	logSugar.Infof("LoadConfig: %v", config)
	return
}
