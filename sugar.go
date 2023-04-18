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
		Level:       LevelToSlogLevel(lg.Level),
		ReplaceAttr: replace,
	}

	var log *slog.Logger
	if lg.Type == "json" {
		log = slog.New(opts.NewJSONHandler(os.Stdout))
	} else {
		log = slog.New(opts.NewTextHandler(os.Stdout))
	}
	return log
}

func New() *slog.Logger {
	lg := &Log{}
	return NewSlog(lg)

}
