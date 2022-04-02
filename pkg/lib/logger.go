package lib

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"kang/global"
	util2 "kang/pkg/util"
	"os"
	"time"
)


// NewLogger 构造日志服务
func NewLogger() (logger *zap.Logger) {
	if ok, _ := util2.PathExists(global.G_Conf.Log.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.G_Conf.Log.Director)
		_ = os.Mkdir(global.G_Conf.Log.Director, os.ModePerm)
	}

	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("./%s/server_debug.log", global.G_Conf.Log.Director), debugPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_info.log",global.G_Conf.Log.Director),infoPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_warn.log",global.G_Conf.Log.Director),warnPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_error.log",global.G_Conf.Log.Director),errorPriority),
	}

	logger = zap.New(zapcore.NewTee(cores[:]...),zap.AddCaller())

	if global.G_Conf.Log.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := util2.GetWriteSyncer(fileName) // 使用file-rotatelogs进行日志分割
	return zapcore.NewCore(getEncoder(), writer, level)
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if global.G_Conf.Log.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderG_Config())
	}
	return zapcore.NewConsoleEncoder(getEncoderG_Config())
}

// getEncoderG_Config 获取zapcore.EncoderG_Config
func getEncoderG_Config() (G_Config zapcore.EncoderConfig) {
	G_Config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.G_Conf.Log.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case global.G_Conf.Log.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		G_Config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.G_Conf.Log.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		G_Config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.G_Conf.Log.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		G_Config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.G_Conf.Log.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		G_Config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		G_Config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return G_Config
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.G_Conf.Log.Prefix + "2006/01/02 - 15:04:05.000"))
}
