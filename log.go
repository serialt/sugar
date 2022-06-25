package sugar

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	LogLevel      string // 日志级别
	LogFile       string // 日志文件存放路径,如果为空，则输出到控制台
	LogType       string // 日志类型，支持 txt 和 json ，默认txt
	LogMaxSize    int    //单位M
	LogMaxBackups int    // 日志文件保留个数
	LogMaxAge     int    // 单位天
	LogCompress   bool   // 压缩轮转的日志
}

func LevelToZapLevel(level string) zapcore.Level {
	// 转换日志级别
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}

}

func (lg *Logger) NewMyLogger() *zap.Logger {

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(LevelToZapLevel(lg.LogLevel))

	// 输出的消息
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		FunctionKey:   zapcore.OmitKey,
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder, // zapcore.CapitalLevelEncoder //按级别显示不同颜色
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05")) //指定时间格式
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 日志输出类型
	var encoder zapcore.Encoder
	switch lg.LogType {
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)

	}

	var core zapcore.Core
	if lg.LogFile == "" {
		core = zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), atomicLevel)
	} else {
		file := zapcore.AddSync(&lumberjack.Logger{
			Filename:   lg.LogFile,       // 日志文件
			MaxSize:    lg.LogMaxSize,    // 单个日志文件大小，单位M
			MaxBackups: lg.LogMaxBackups, // 轮转保留个数
			MaxAge:     lg.LogMaxAge,     // 最长保留时间，单位天
			Compress:   lg.LogCompress,   // 默认不压缩
		})
		core = zapcore.NewCore(encoder, zapcore.AddSync(file), atomicLevel)
	}

	// 开启开发模式，堆栈跟踪: [zap.AddCaller()]
	myLogger := zap.New(core, zap.AddCaller())
	return myLogger
}

// NewLogger 自定日志配置可以参考此方法
func NewLogger(logLevel, logFile string) *zap.Logger {
	lg := &Logger{
		LogLevel:      logLevel,
		LogFile:       logFile,
		LogType:       "txt",
		LogMaxSize:    50,
		LogMaxBackups: 10,
		LogMaxAge:     365,
		LogCompress:   true,
	}
	return lg.NewMyLogger()
}

// NewSugarLogger 创建一个sugar
func NewSugarLogger(logLevel, logFile string) *zap.SugaredLogger {
	sugarLog := NewLogger(logLevel, logFile)
	return sugarLog.Sugar()

}
