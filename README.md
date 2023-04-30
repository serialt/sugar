# sugar lib

封装好的一些常用方法

### 使用方法

```
go get github.com/serialt/sugar/v2

go get golang.org/x/exp/slog

```

### 库使用方法

#### 简单使用日志

```go
package main

import (
	"github.com/serialt/sugar/v2"
	"golang.org/x/exp/slog"
)

func init() {
	slog.SetDefault(sugar.New())
}

func main() {
	slog.Debug("debug", "hello", "world")
	slog.Info("info", "hello", "world")
	slog.Error("error", "hello", "world")
}

```

#### 可选参数

```go
package main

import (
	sugar "github.com/serialt/sugar/v2"
	"golang.org/x/exp/slog"
)

func init() {
	options := []sugar.LogOptions{
		sugar.WithLevel("debug"),
		sugar.WithShort(false),
		sugar.WithType("json"),
	}
	slog.SetDefault(sugar.New(options...))
}
func main() {

	slog.Debug("debug msg")
	slog.Info("info msg")
}

```


#### 复杂使用
```go
package main

import (
	sugar "github.com/serialt/sugar/v2"
	flag "github.com/spf13/pflag"
	"golang.org/x/exp/slog"
)

type Config struct {
	Server   string
	Port     string
	LogLevel string
}

var (
	cfgfile string
	config  *Config
)

func init() {

	flag.StringVarP(&cfgfile, "config", "c", "config.yaml", "Config file")
	flag.CommandLine.SortFlags = false
	flag.Parse()

	err := sugar.LoadConfig(cfgfile, &config)
	if err != nil {
		config = new(Config)
	}
	options := []sugar.LogOptions{
		sugar.WithLevel(config.LogLevel),
	}
	slog.SetDefault(sugar.New(options...))
}
func main() {
	slog.Info("struct", "cfg", config)
	slog.Debug("debug msg")
	slog.Info("info msg")
}

```

