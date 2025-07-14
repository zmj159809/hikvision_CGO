package logger

import (
	"os"
	"path/filepath"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var (
	// Logger 全局日志实例
	Logger *zap.Logger
	// Sugar 便于使用的 sugar logger
	Sugar *zap.SugaredLogger
	// mutex 用于线程安全的初始化
	mutex sync.RWMutex
	// initialized 标记是否已初始化
	initialized bool
)

// Init 初始化日志系统
func Init(config *Config) error {
	mutex.Lock()
	defer mutex.Unlock()

	if config == nil {
		config = DefaultConfig()
	}

	// 创建日志目录
	if err := os.MkdirAll(filepath.Dir(config.LogFilePath), 0755); err != nil {
		return err
	}

	// 配置编码器
	var encoderConfig zapcore.EncoderConfig
	if config.Development {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
	}

	// 自定义时间格式
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	// 创建 Core
	var cores []zapcore.Core

	// 文件输出
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.LogFilePath,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	})
	fileCore := zapcore.NewCore(fileEncoder, zap.CombineWriteSyncers(fileWriter, os.Stdout), config.Level)
	cores = append(cores, fileCore)

	// 控制台输出
	if config.EnableConsole {
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), config.Level)
		cores = append(cores, consoleCore)
	}

	// 创建 logger
	core := zapcore.NewTee(cores...)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	if config.Development {
		logger = logger.WithOptions(zap.Development())
	}

	Logger = logger
	Sugar = logger.Sugar()
	initialized = true

	return nil
}

// IsInitialized 检查日志系统是否已初始化
func IsInitialized() bool {
	mutex.RLock()
	defer mutex.RUnlock()
	return initialized
}

// Close 关闭日志系统
func Close() {
	mutex.Lock()
	defer mutex.Unlock()

	if Logger != nil {
		Logger.Sync()
		Logger = nil
	}
	if Sugar != nil {
		Sugar.Sync()
		Sugar = nil
	}
	initialized = false
}

// ensureInitialized 确保日志系统已初始化，如果未初始化则使用默认配置
func ensureInitialized() {
	if !IsInitialized() {
		Init(DefaultConfig())
	}
}

// Info 记录 Info 级别日志
func Info(msg string, fields ...zap.Field) {
	ensureInitialized()
	if Logger != nil {
		Logger.Info(msg, fields...)
	}
}

// Error 记录 Error 级别日志
func Error(msg string, fields ...zap.Field) {
	ensureInitialized()
	if Logger != nil {
		Logger.Error(msg, fields...)
	}
}

// Warn 记录 Warn 级别日志
func Warn(msg string, fields ...zap.Field) {
	ensureInitialized()
	if Logger != nil {
		Logger.Warn(msg, fields...)
	}
}

// Debug 记录 Debug 级别日志
func Debug(msg string, fields ...zap.Field) {
	ensureInitialized()
	if Logger != nil {
		Logger.Debug(msg, fields...)
	}
}

// Fatal 记录 Fatal 级别日志
func Fatal(msg string, fields ...zap.Field) {
	ensureInitialized()
	if Logger != nil {
		Logger.Fatal(msg, fields...)
	}
}

// Panic 记录 Panic 级别日志
func Panic(msg string, fields ...zap.Field) {
	ensureInitialized()
	if Logger != nil {
		Logger.Panic(msg, fields...)
	}
}

// Infof 使用格式化字符串记录 Info 级别日志
func Infof(template string, args ...interface{}) {
	ensureInitialized()
	if Sugar != nil {
		Sugar.Infof(template, args...)
	}
}

// Errorf 使用格式化字符串记录 Error 级别日志
func Errorf(template string, args ...interface{}) {
	ensureInitialized()
	if Sugar != nil {
		Sugar.Errorf(template, args...)
	}
}

// Warnf 使用格式化字符串记录 Warn 级别日志
func Warnf(template string, args ...interface{}) {
	ensureInitialized()
	if Sugar != nil {
		Sugar.Warnf(template, args...)
	}
}

// Debugf 使用格式化字符串记录 Debug 级别日志
func Debugf(template string, args ...interface{}) {
	ensureInitialized()
	if Sugar != nil {
		Sugar.Debugf(template, args...)
	}
}

// Fatalf 使用格式化字符串记录 Fatal 级别日志
func Fatalf(template string, args ...interface{}) {
	ensureInitialized()
	if Sugar != nil {
		Sugar.Fatalf(template, args...)
	}
}

// Panicf 使用格式化字符串记录 Panic 级别日志
func Panicf(template string, args ...interface{}) {
	ensureInitialized()
	if Sugar != nil {
		Sugar.Panicf(template, args...)
	}
}

// InfoArgs 使用 Sugar Logger 记录 Info 级别日志
func InfoArgs(args ...interface{}) {
	ensureInitialized()
	if Sugar != nil {
		Sugar.Info(args...)
	}
}

// ErrorArgs 使用 Sugar Logger 记录 Error 级别日志
func ErrorArgs(args ...interface{}) {
	ensureInitialized()
	if Sugar != nil {
		Sugar.Error(args...)
	}
}

// WarnArgs 使用 Sugar Logger 记录 Warn 级别日志
func WarnArgs(args ...interface{}) {
	ensureInitialized()
	if Sugar != nil {
		Sugar.Warn(args...)
	}
}

// DebugArgs 使用 Sugar Logger 记录 Debug 级别日志
func DebugArgs(args ...interface{}) {
	ensureInitialized()
	if Sugar != nil {
		Sugar.Debug(args...)
	}
}

// With 添加字段到日志上下文
func With(fields ...zap.Field) *zap.Logger {
	ensureInitialized()
	if Logger != nil {
		return Logger.With(fields...)
	}
	return nil
}

// WithOptions 添加选项到日志上下文
func WithOptions(opts ...zap.Option) *zap.Logger {
	ensureInitialized()
	if Logger != nil {
		return Logger.WithOptions(opts...)
	}
	return nil
}

// Named 创建一个命名的子logger
func Named(name string) *zap.Logger {
	ensureInitialized()
	if Logger != nil {
		return Logger.Named(name)
	}
	return nil
}
