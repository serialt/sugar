# 连接数据库和日志的库

### 使用方法
```
go get -u  github.com/serialt/sugar

```

### 库使用方法
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
	// 日志配置
	LogLevel      = "info"
	LogFile       = ""    // 日志文件存放路径,如果为空，则输出到控制台
	LogType       = "txt" // 日志类型，支持 txt 和 json ，默认txt
	LogMaxSize    = 100   //单位M
	LogMaxBackups = 3     // 日志文件保留个数
	LogMaxAge     = 365   // 单位天
	LogCompress   = true  // 压缩轮转的日志

	// 版本信息
	APPName    = ""
	APPVersion = ""

	Logger   *zap.Logger
	LogSugar *zap.SugaredLogger
	// DB       *gorm.DB
)

func env(key, def string) string {
	if x := os.Getenv(key); x != "" {
		return x
	}
	return def
}

func init() {
	flag.StringVar(&LogFile, "logFile", env("LogFile", LogFile), "Logfile path")
	flag.StringVar(&LogLevel, "logLevel", env("LogLevel", LogLevel), "Log level, debug,info,warn,error,panic")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Println("使用说明")
		flag.PrintDefaults()
	}
	flag.ErrHelp = fmt.Errorf("\n\nSome errors have occurred, check and try again !!! ")

	flag.CommandLine.SortFlags = false
	flag.Parse()

	mylg := &sugar.Logger{
		LogLevel:      LogLevel,
		LogFile:       LogFile,
		LogType:       LogType,
		LogMaxSize:    LogMaxSize,
		LogMaxBackups: LogMaxBackups,
		LogMaxAge:     LogMaxAge,
		LogCompress:   LogCompress,
	}
	Logger = mylg.NewMyLogger()
	LogSugar = Logger.Sugar()

	// mydb := &sugar.Database{
	// 	Type:     "mysql",
	// 	Addr:     "host",
	// 	Port:     "3306",
	// 	DBName:   "db-name",
	// 	Username: "db-user",
	// 	Password: "db-pass",
	// }
	// DB = mydb.NewDBConnect(Logger)
	// DB.AutoMigrate(&Department{})
	// DB.AutoMigrate(&Userlist{})

}

func main() {
	LogSugar.Debug("debug log")
	LogSugar.Info("info log")
	LogSugar.Error("error log")
}


```


配置文件示例：
```go
var (
	// 版本信息
	APPName    = ""
	Maintainer = ""
	APPVersion = ""
	BuildTime  = ""
	GitCommit  = ""
	GOVERSION  = runtime.Version()
	GOOSARCH   = runtime.GOOS + "/" + runtime.GOARCH
	// 其他配置文件
	ConfigPath = ""
	DefaultConfigFile = ".git-audit.yaml"

	Logger   *zap.Logger
	LogSugar *zap.SugaredLogger
	DB       *gorm.DB
)

type Database struct {
	Type            string        `yaml:"type"`
	DBName          string        `gorm:"dbname" yaml:"dbname"`
	Addr            string        `gorm:"addr" yaml:"addr"`
	Port            string        `gorm:"port" yaml:"port"`
	Username        string        `gorm:"username" yaml:"username"`
	Password        string        `gorm:"password" yaml:"password"`
	MaxOpenConns    int           `yaml:"maxOpenConns"`
	MaxIdleConns    int           `yaml:"maxIdleConns"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"`
}

type GitLog struct {
	LogLevel string `yaml:"logLevel"` // 日志级别
	LogFile  string `yaml:"logFile"`  // 日志文件存放路径,如果为空，则输出到控制台
	LogType  string `yaml:"logType"`  // 日志类型，支持 txt 和 json ，默认txt
	// LogMaxSize    int    //单位M
	// LogMaxBackups int    // 日志文件保留个数
	// LogMaxAge     int    // 单位天
	// LogCompress   bool   // 压缩轮转的日志
}

type MyConfig struct {
	Database Database `json:"database" yaml:"database"`
	GitLog   GitLog   `yaml:"gitLog"`
}

var Config *MyConfig


var Config *MyConfig

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

func LoadConfig(filepath string) {
	if filepath == "" {
		dir, _ := homedir.Dir()
		filepath = fmt.Sprintf("%v/%v", dir, DefaultConfigFile)
	}
	filepath, err := homedir.Expand(filepath)
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
	err = yaml.Unmarshal(config, &Config)
	if err != nil {
		fmt.Printf("Unmarshal to struct, err: %v", err)
	}
	// fmt.Printf("LoadConfig: %v\n", Config)
	fmt.Printf("Config path: %v\n", filepath)
}

```