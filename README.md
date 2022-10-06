# sugar lib
封装好的一些常用方法

### 使用方法
```
go get github.com/serialt/sugar

```

### 库使用方法

#### 简单使用日志
```go
package main

func main() {
	// 设置简单日志参数
	sugar.SetLog("error", "log.txt")
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

	// Logger   *zap.Logger
	LogSugar *zap.SugaredLogger
)


func init() {
	
	// Logger = sugar.NewLogger(LogLevel, LogFile)
	LogSugar = sugar.NewSugarLogger("info", "ccc", "", false)
}

func main() {
	LogSugar.Debug("debug log")
	LogSugar.Info("info log")
	LogSugar.Error("error log")
}
```
更多复杂使用可以参考sugar.NewLogger的实现


```go
package main

import (
	"fmt"
	"os"

	"github.com/serialt/sugar"
	flag "github.com/spf13/pflag"
	"go.uber.org/zap"
)

var (
	// 版本信息
	appVersion bool // 控制是否显示版本
	APPVersion = "v0.0.2"
	BuildTime  = "2006-01-02 15:04:05"
	GitCommit  = "xxxxxxxxxxx"

	// 配置文件,置空则表示读取项目根目录里的config.yaml
	ConfigPath = "config.yaml"

	// Logger   *zap.Logger
	LogSugar *zap.SugaredLogger
	Config   *MyConfig
)

type Log struct {
	LogLevel string `yaml:"logLevel"` // 日志级别，支持debug,info,warn,error,panic
	LogFile  string `yaml:"logFile"`  // 日志文件存放路径,如果为空，则输出到控制台
}

type MyConfig struct {
	Log Log    `yaml:"log"`
	Msg string `yaml:"msg"`
}

func init() {
	flag.BoolVarP(&appVersion, "version", "v", false, "Display build and version msg")
	flag.StringVarP(&ConfigPath, "cfgFile", "c", ConfigPath, "")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Println("使用说明")
		flag.PrintDefaults()
	}
	flag.ErrHelp = fmt.Errorf("\n\nSome errors have occurred, check and try again !!! ")
	flag.CommandLine.SortFlags = false
	flag.Parse()

	// 读取配置文件
	err := sugar.LoadConfig(ConfigPath, &Config)
	if err != nil {
		Config = new(MyConfig)
	}
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