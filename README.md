# 连接数据库和日志的库

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

