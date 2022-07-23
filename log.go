package sugar

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var sugar *zap.SugaredLogger

type Logger struct {
	LogLevel      string // 日志级别
	LogFile       string // 日志文件存放路径,如果为空，则输出到控制台
	LogType       string // 日志类型，支持 txt 和 json ，默认txt
	LogMaxSize    int    //单位M
	LogMaxBackups int    // 日志文件保留个数
	LogMaxAge     int    // 单位天
	LogCompress   bool   // 压缩轮转的日志
	LogColor      bool   // 日志级别分颜色
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
	// case "dpanic":
	// 	return zapcore.DPanicLevel
	// case "panic":
	// 	return zapcore.PanicLevel
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
		//  zapcore.LowercaseLevelEncoder 小写日志级别无颜色
		//  zapcore.LowercaseColorLevelEncoder 小写日志级别无颜色
		//  zapcore.CapitalLevelEncoder
		//  zapcore.CapitalColorLevelEncoder
		EncodeLevel: zapcore.CapitalLevelEncoder,
		// EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		// 	enc.AppendString(t.Format("2006-01-02 15:04:05 Mon")) //指定时间格式
		// },
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行消耗的时间转化成浮点型的秒
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 以包/文件:行号 格式化调用堆栈
	}
	// 简易设置日志级别分颜色
	if lg.LogColor {
		// capital, capitalColor, color, default LowercaseLevelEncoder
		encoderConfig.EncodeLevel.UnmarshalText([]byte("capitalColor"))
	}
	// 简易设置日志时间
	// // RFC3339Nano 2022-07-23T10:49:00.656748+08:00
	// // RFC3339     2022-07-23T10:49:47+08:00
	// // ISO8601     2022-07-23T10:50:14.760+0800
	// // millis      1.6585446469065e+12
	// // nanos       1658544668921929000
	// encoderConfig.EncodeTime.UnmarshalText([]byte("RFC3339"))

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
func NewLogger(logLevel, logFile, logtype string, logColor bool) *zap.Logger {
	lg := &Logger{
		LogLevel:      logLevel,
		LogFile:       logFile,
		LogType:       logtype,
		LogMaxSize:    50,
		LogMaxBackups: 10,
		LogMaxAge:     365,
		LogCompress:   true,
		LogColor:      logColor,
	}
	return lg.NewMyLogger()
}

func init() {
	sugar = NewLogger("info", "", "", false).Sugar()
}

// NewSugarLogger 创建一个sugar
func NewSugarLogger(logLevel, logFile, logtype string, logColor bool) *zap.SugaredLogger {
	sugarLog := NewLogger(logLevel, logFile, logtype, logColor)
	return sugarLog.Sugar()

}

func Debug(args ...interface{}) {
	sugar.Debug(args...)
}

func Debugf(temp string, args ...interface{}) {
	sugar.Debugf(temp, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	sugar.Debugw(msg, keysAndValues...)
}

func Info(args ...interface{}) {
	sugar.Info(args...)
}

func Infof(temp string, args ...interface{}) {
	sugar.Infof(temp, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	sugar.Infow(msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	sugar.Warn(args...)
}

func Warnf(temp string, args ...interface{}) {
	sugar.Warnf(temp, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	sugar.Warnw(msg, keysAndValues...)
}

func Error(args ...interface{}) {
	sugar.Error(args...)
}

func Errorf(temp string, args ...interface{}) {
	sugar.Errorf(temp, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	sugar.Errorw(msg, keysAndValues...)
}

func Fatal(args ...interface{}) {
	sugar.Fatal(args...)
}

func Fatalf(temp string, args ...interface{}) {
	sugar.Fatalf(temp, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	sugar.Fatalw(msg, keysAndValues...)
}
