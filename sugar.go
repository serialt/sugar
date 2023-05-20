package sugar

import (
	"io"
	"os"
	"path/filepath"

	"golang.org/x/exp/slog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Log struct {
	Level      string // 日志级别
	File       string // 日志文件存放路径,如果为空，则输出到控制台
	Type       string // 日志类型，支持 txt 和 json ，默认txt
	MaxSize    int    //单位M
	MaxBackups int    // 日志文件保留个数
	MaxAge     int    // 单位天
	Compress   bool   // 压缩轮转的日志
	Short      bool   // 以包/文件:行号 显示短路径，不显示全路径
}

type LogOptions func(*Log)

func WithLevel(level string) LogOptions {
	return func(lg *Log) {
		lg.Level = level
	}
}
func WithFile(file string) LogOptions {
	return func(lg *Log) {
		lg.File = file
	}
}

func WithType(tp string) LogOptions {
	return func(lg *Log) {
		lg.Type = tp
	}
}

func WithMaxSize(maxSize int) LogOptions {
	return func(lg *Log) {
		lg.MaxSize = maxSize
	}
}

func WithCompress(compress bool) LogOptions {
	return func(lg *Log) {
		lg.Compress = compress
	}
}

func WithShort(Short bool) LogOptions {
	return func(lg *Log) {
		lg.Short = Short
	}
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
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		return a
	}
	opts := &slog.HandlerOptions{
		AddSource:   true,
		Level:       LevelToSlogLevel(lg.Level),
		ReplaceAttr: replace,
	}
	var out io.Writer
	if len(lg.File) > 0 {
		out = &lumberjack.Logger{
			Filename:   "slog.log",    // 日志文件
			MaxSize:    lg.MaxSize,    // 单个日志文件大小，单位M
			MaxBackups: lg.MaxBackups, // 轮转保留个数
			MaxAge:     lg.MaxAge,     // 最长保留时间，单位天
			Compress:   lg.Compress,   // 默认不压缩
		}
	} else {
		out = os.Stdout
	}

	var log *slog.Logger
	if lg.Type == "json" {
		log = slog.New(slog.NewJSONHandler(out, opts))
	} else {
		log = slog.New(slog.NewTextHandler(out, opts))
	}
	return log
}

func New(options ...LogOptions) *slog.Logger {
	// 默认值的设定
	lg := &Log{
		Level:      "info",
		Type:       "txt",
		MaxSize:    100,
		MaxBackups: 30,
		MaxAge:     365,
		Compress:   true,
		Short:      true,
	}

	// 遍历可选参数，然后分别调用匿名函数，将连接对象指针传入，进行修改
	for _, op := range options {
		// 遍历调用函数，进行数据修改
		op(lg)
	}
	return NewSlog(lg)
}

// // SetLog 用于配置简单的日志
// func SetLog(level string, file string) {
// 	std.LogLevel = level
// 	std.LogFile = file
// 	sugar = std.newInnerLogger(std.NewCore()).Sugar()
// 	Log = std.newInnerLogger(std.NewCore())

// }
