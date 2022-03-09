package sugar

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

type Database struct {
	Type            string
	Addr            string
	Port            string
	DBName          string
	Username        string
	Password        string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

var logSugar *zap.SugaredLogger

func GetMysqlGormDB(mydb *Database) *gorm.DB {
	var myGormDB *gorm.DB
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mydb.Username,
		mydb.Password,
		mydb.Addr,
		mydb.Port,
	)
	for {
		myGormDB, err = gorm.Open(mysql.New(mysql.Config{
			DSN:                       dsn,
			DefaultStringSize:         256,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		}), &gorm.Config{Logger: GormLogger})
		if err != nil {
			logSugar.Infof("grom open failed: %v\n", err)
		} else {
			break
		}
	}
	return myGormDB
}

func GetPostgreSQLGormDB(mydb *Database) *gorm.DB {
	var myGormDB *gorm.DB
	var err error
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		mydb.Username,
		mydb.Password,
		mydb.Addr,
		mydb.Port,
	)
	for {
		myGormDB, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage,
		}), &gorm.Config{Logger: GormLogger})
		if err != nil {
			logSugar.Infof("gorm open failed: %v\n", err)
		} else {
			break
		}
	}
	return myGormDB
}

// 设置gorm日志使用zap
var GormLogger zapgorm2.Logger

// 创建mydb
func (db *Database) NewDBConnect(zaplog *zap.Logger) *gorm.DB {
	// 使用zap 接收gorm日志
	GormLogger = zapgorm2.New(zaplog)
	GormLogger.SetAsDefault()
	logSugar = zaplog.Sugar()

	var GormDB *gorm.DB
	switch db.Type {
	// case "sqlite":
	// 	mySqlite := NewMySqlite()
	// 	GormDB = SqliteInit(mySqlite)
	// 	fmt.Println("使用的数据库是 sqlite")
	case "mysql":
		GormDB = GetMysqlGormDB(db)
		logSugar.Info("使用的数据库是 mysql")
	case "postgresql":
		GormDB = GetPostgreSQLGormDB(db)
		logSugar.Info("使用的数据库是 postgresql")
	default:
		logSugar.Info("The database is not supported, please choice [sqlite] or [mysql]")
	}
	if db.MaxOpenConns != 0 {

		// Gorm 使用database/sql 维护连接池
		sqlDB, _ := GormDB.DB()

		// 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(db.MaxIdleConns)

		// 设置打开数据库连接的最大数量
		sqlDB.SetMaxOpenConns(db.MaxOpenConns)

		// 设置了连接可复用的最大时间
		sqlDB.SetConnMaxLifetime(db.ConnMaxLifetime)
	} else {
		sqlDB, _ := GormDB.DB()

		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Hour)

	}

	return GormDB

}
