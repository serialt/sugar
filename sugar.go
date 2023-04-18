package sugar

import (
	"io"
	"os"
	"path/filepath"

	"golang.org/x/exp/slog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Log struct {
	LogLevel      string // 日志级别
	LogFile       string // 日志文件存放路径,如果为空，则输出到控制台
	LogType       string // 日志类型，支持 txt 和 json ，默认txt
	LogMaxSize    int    //单位M
	LogMaxBackups int    // 日志文件保留个数
	LogMaxAge     int    // 单位天
	LogCompress   bool   // 压缩轮转的日志
	LogColor      bool   // 日志级别分颜色
	Short         bool   // 以包/文件:行号 显示短路径，不显示全路径
}

// LevelToZapLevel  转换日志级别
func LevelToSlogLevel(level string) slog.Level {
	switch level {
	case "debug", "DEBUG":
		return slog.LevelDebug
	case "info", "INFO":
		return slog.LevelInfo
	case "warn", "WARN", "WARNING":
		return slog.LevelWarn
	case "error", "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}

}

func NewSlog(lg *Log) *slog.Logger {

	replace := func(groups []string, a slog.Attr) slog.Attr {

		// Remove the directory from the source's filename.
		if a.Key == slog.SourceKey && lg.Short {
			a.Value = slog.StringValue(filepath.Base(a.Value.String()))
		}
		return a
	}
	opts := slog.HandlerOptions{
		AddSource:   true,
		Level:       LevelToSlogLevel(lg.LogLevel),
		ReplaceAttr: replace,
	}
	var out io.Writer
	if len(lg.LogFile) > 0 {
		out = &lumberjack.Logger{
			Filename:   "slog.log", // 日志文件
			MaxSize:    100,        // 单个日志文件大小，单位M
			MaxBackups: 30,         // 轮转保留个数
			MaxAge:     365,        // 最长保留时间，单位天
			Compress:   true,       // 默认不压缩
		}
	} else {
		out = os.Stdout
	}

	var log *slog.Logger
	if lg.LogType == "json" {
		log = slog.New(opts.NewJSONHandler(out))
	} else {
		log = slog.New(opts.NewTextHandler(out))
	}
	return log
}

func New() *slog.Logger {
	lg := &Log{}
	return NewSlog(lg)

}

// // SetLog 用于配置简单的日志
// func SetLog(level string, file string) {
// 	std.LogLevel = level
// 	std.LogFile = file
// 	sugar = std.newInnerLogger(std.NewCore()).Sugar()
// 	Log = std.newInnerLogger(std.NewCore())

// }
