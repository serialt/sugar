package sugar

import (
	"os"
	"path/filepath"

	"golang.org/x/exp/slog"
)

type Log struct {
	Level string // 日志级别  string // 日志文件存放路径,如果为空，则输出到控制台
	Type  string // 日志类型，支持 txt 和 json ，默认txt
	Short bool   // 以包/文件:行号 显示短路径，不显示全路径
}

type LogOptions func(*Log)

func WithLevel(level string) LogOptions {
	return func(lg *Log) {
		lg.Level = level
	}
}

func WithType(tp string) LogOptions {
	return func(lg *Log) {
		lg.Type = tp
	}
}

func WithShort(short bool) LogOptions {
	return func(lg *Log) {
		lg.Short = short
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
		// if a.Key == slog.TimeKey {
		// 	a.Value = slog.AnyValue(time.Now().Format(time.RFC3339))
		// }

		// Remove the directory from the source's filename.
		if a.Key == slog.SourceKey && lg.Short {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
			// a.Value = slog.StringValue(filepath.Base(a.Value.String()))
		}
		return a
	}
	opts := &slog.HandlerOptions{
		AddSource:   true,
		Level:       LevelToSlogLevel(lg.Level),
		ReplaceAttr: replace,
	}

	var log *slog.Logger
	if lg.Type == "json" {
		log = slog.New(slog.NewJSONHandler(os.Stderr, opts))
	} else {
		log = slog.New(slog.NewTextHandler(os.Stdout, opts))
	}
	return log
}

func New(options ...LogOptions) *slog.Logger {
	// 默认值的设定
	lg := &Log{
		Level: "info",
		Type:  "txt",
		Short: true,
	}

	// 遍历可选参数，然后分别调用匿名函数，将连接对象指针传入，进行修改
	for _, opt := range options {
		// 遍历调用函数，进行数据修改
		opt(lg)
	}
	return NewSlog(lg)
}
