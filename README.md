# 连接数据库和日志的库

### 使用方法
```
go get -u  github.com/serialt/sugar

```

### 库使用方法
```go
package main

import (
	"github.com/serialt/sugar"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger   *zap.Logger
	LogSugar *zap.SugaredLogger
	DB       *gorm.DB
)

func init() {
	mylg := &sugar.Logger{
		LogLevel:      "debug",
		LogFile:       "",
		LogType:       "txt",
		LogMaxSize:    50,
		LogMaxBackups: 3,
		LogMaxAge:     365,
		LogCompress:   true,
	}
	Logger = mylg.NewMyLogger()
	LogSugar = Logger.Sugar()

	mydb := &sugar.Database{
		Type:     "mysql",
		Addr:     "host",
		Port:     "3306",
		DBName:   "db-name",
		Username: "db-user",
		Password: "db-pass",
	}
	DB = mydb.NewDBConnect(Logger)
	// DB.AutoMigrate(&Department{})
	// DB.AutoMigrate(&Userlist{})


}

```