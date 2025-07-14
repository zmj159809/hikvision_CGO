package logger

import (
	"go.uber.org/zap/zapcore"
)

// Config 日志配置
type Config struct {
	Level         zapcore.Level // 日志级别
	LogFilePath   string        // 日志文件路径
	MaxSize       int           // 单个日志文件最大大小 (MB)
	Compress      bool          // 是否启用压缩
	MaxBackups    int           // 保留的旧日志文件数量
	MaxAge        int           // 日志文件保留天数
	Development   bool          // 是否为开发模式
	EnableConsole bool          // 是否启用控制台输出
}

// DefaultConfig 默认日志配置
func DefaultConfig() *Config {
	return &Config{
		Compress:      true,
		Level:         zapcore.InfoLevel,
		LogFilePath:   "./logs/hikvision_cgo.log",
		MaxSize:       100,
		MaxBackups:    3,
		MaxAge:        30,
		Development:   false,
		EnableConsole: true,
	}
}

// DevelopmentConfig 开发模式日志配置
func DevelopmentConfig() *Config {
	return &Config{
		Compress:      false,
		Level:         zapcore.DebugLevel,
		LogFilePath:   "./logs/hikvision_cgo_dev.log",
		MaxSize:       50,
		MaxBackups:    5,
		MaxAge:        7,
		Development:   true,
		EnableConsole: true,
	}
}

// ProductionConfig 生产模式日志配置
func ProductionConfig() *Config {
	return &Config{
		Compress:      true,
		Level:         zapcore.InfoLevel,
		LogFilePath:   "./logs/hikvision_cgo_prod.log",
		MaxSize:       200,
		MaxBackups:    10,
		MaxAge:        90,
		Development:   false,
		EnableConsole: false,
	}
}
