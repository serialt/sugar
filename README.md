# 连接数据库和日志的库

### 使用方法
v1及以下版本是带有数据库，v2版本不带有数据库
```
go get -u  github.com/serialt/sugar/v2

```

### 库使用方法

#### 简单使用日志
```go
package main

func main() {
	// 默认日志配置是: info级别，输出到控制台
	sugar.Debug("debug logSugar")
	sugar.Info("info logSugar")
	sugar.Error("error logSugar")
}
```

#### 复杂使用
```go
package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/serialt/sugar"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

var (
	// 版本信息
	appVersion bool // 控制是否显示版本
	APPVersion = "v0.0.2"
	BuildTime  = "2006-01-02 15:04:05"
	GitCommit  = "xxxxxxxxxxx"

	// 配置文件,置空则表示读取项目根目录里的config.yaml
	ConfigPath = ""

	// Logger   *zap.Logger
	LogSugar *zap.SugaredLogger
	Config   *MyConfig
)

type Gitee struct {
	User    string `yaml:"user"`
	Private string `yaml:"private"`
}

type Log struct {
	LogLevel string `yaml:"logLevel"` // 日志级别，支持debug,info,warn,error,panic
	LogFile  string `yaml:"logFile"`  // 日志文件存放路径,如果为空，则输出到控制台
}
type MyConfig struct {
	Log   Log   `yaml:"log"`
	Gitee Gitee `yaml:"gitee"`
}

func init() {
	flag.BoolVarP(&appVersion, "version", "v", false, "Display build and version msg")
	flag.StringVarP(&ConfigPath, "cfgFile", "c", sugar.Env("CONFIG", ConfigPath), "")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Println("使用说明")
		flag.PrintDefaults()
	}
	flag.ErrHelp = fmt.Errorf("\n\nSome errors have occurred, check and try again !!! ")
	flag.CommandLine.SortFlags = false
	flag.Parse()

	// 读取配置文件
	byteConfig, _ := sugar.LoadConfig(ConfigPath)
	yaml.Unmarshal(byteConfig, &Config)
	// fmt.Println(Config)

	// Logger = sugar.NewLogger(LogLevel, LogFile)
	LogSugar = sugar.NewSugarLogger(Config.Log.LogLevel, Config.Log.LogFile, "", false)
}

func main() {
	if appVersion {
		fmt.Printf("APPVersion: %v  BuildTime: %v  GitCommit: %v\n",
			APPVersion,
			BuildTime,
			GitCommit)
		return
	}
	LogSugar.Debug("debug log")
	LogSugar.Info("info log")
	LogSugar.Error("error log")
}
```

配置文件示例
```yaml
log:
  logLevel: info
  # logFile: imau.log
gitee:
 user: imaus
 private: true 

```