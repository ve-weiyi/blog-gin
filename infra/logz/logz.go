package logz

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 日志级别取值
const (
	levelDebug = "debug"
	levelInfo  = "info"
	levelWarn  = "warn"
	levelError = "error"
)

// 输出模式取值，可组合使用（如 "console|file"）
const (
	modeConsole = "console"
	modeFile    = "file"
)

// Config 日志配置。
type Config struct {
	Level      string `json:"level"`
	Mode       string `json:"mode"` // console|file
	Filename   string `json:"file_name"`
	MaxSize    int    `json:"max_size"`
	MaxBackups int    `json:"max_backups"`
	MaxAge     int    `json:"max_age"`
	Compress   bool   `json:"compress"`
}

// New 按配置创建 logger。
// Mode 同时包含 console 与 file 时，日志同时写入标准输出与文件。
func New(conf *Config, opts ...zap.Option) (*zap.Logger, error) {
	level, err := parseLevel(conf.Level)
	if err != nil {
		return nil, err
	}

	var cores []zapcore.Core

	if strings.Contains(conf.Mode, modeConsole) {
		consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig())
		cores = append(cores, zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level))
	}

	if strings.Contains(conf.Mode, modeFile) {
		if err := os.MkdirAll(filepath.Dir(conf.Filename), 0755); err != nil {
			return nil, fmt.Errorf("logz: create log dir: %w", err)
		}

		hook := lumberjack.Logger{
			Filename:   conf.Filename,
			MaxSize:    conf.MaxSize,
			MaxBackups: conf.MaxBackups,
			MaxAge:     conf.MaxAge,
			Compress:   conf.Compress,
		}
		fileEncoder := zapcore.NewJSONEncoder(jsonEncoderConfig())
		cores = append(cores, zapcore.NewCore(fileEncoder, zapcore.AddSync(&hook), level))
	}

	if len(cores) == 0 {
		return nil, fmt.Errorf("logz: mode %q matches neither %s nor %s", conf.Mode, modeConsole, modeFile)
	}

	return zap.New(zapcore.NewTee(cores...), append(opts, zap.AddCaller())...), nil
}

// defaultLogger 进程级 logger，由 SetLogger 写入、L()/S() 读取。
// 用 atomic.Pointer 持有：SetLogger 可在运行时被任意 goroutine 调用，直接读写裸指针是数据竞争。
var defaultLogger atomic.Pointer[zap.Logger]

// SetLogger 设置进程级 logger，供 L() / S() 取用。
func SetLogger(logger *zap.Logger) {
	defaultLogger.Store(logger)
}

// L 返回进程级 logger。
// 尚未 SetLogger 时返回丢弃全部日志的 nop logger，保证调用方无需判空。
func L() *zap.Logger {
	if l := defaultLogger.Load(); l != nil {
		return l
	}
	return zap.NewNop()
}

// S 返回进程级 SugaredLogger。
func S() *zap.SugaredLogger {
	return L().Sugar()
}

// parseLevel 把配置中的级别字符串转为 zap 级别。
// 空值按 info 处理，其余无法识别的取值返回 error（不静默回落，避免配置写错却毫无察觉）。
func parseLevel(level string) (zapcore.Level, error) {
	switch level {
	case levelDebug:
		return zap.DebugLevel, nil
	case levelInfo, "":
		return zap.InfoLevel, nil
	case levelWarn:
		return zap.WarnLevel, nil
	case levelError:
		return zap.ErrorLevel, nil
	default:
		return zap.InfoLevel, fmt.Errorf("logz: unknown level %q", level)
	}
}

// jsonEncoderConfig 文件输出使用的 JSON 编码配置
func jsonEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
}

// consoleEncoderConfig 控制台输出使用的文本编码配置（带颜色）
func consoleEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("15:04:05"),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
}
