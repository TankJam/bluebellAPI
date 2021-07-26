package logger

import (
	"bluebellAPI/settings"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// 定义全局的 zap 日志结构体对象
var lg *zap.Logger


// Init 初始化lg
func Init(cfg *settings.LogConfig, mode string)(err error){
	fmt.Println(cfg, mode)
	// 1.得到日志写入器
	writeSyncer := getLogWrite(
		cfg.FileName, cfg.MaxAge, cfg.MaxBackups, cfg.MaxSize)

	// 2.获取编码转换之后的对象
	encoder := getEncoder()
	// 3.反序列化日志级别格式化
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil{
		return
	}
	// 4.定义快速的日志记录器 对象
	var core zapcore.Core

	// 5.进入开发者模式，日志输出到终端
	if mode == "dev"{
		// 设置开发者模式，配置终端输出，并返回终端输出对象
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

		// 日志写入组，写入日志文件与输出终端
		core = zapcore.NewTee(
			// NewCore创建一个Core，将日志写入WriteSyncer。
			zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	}else{  // 不是开发者模式，则不输出到终端，直接写入文件中
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}

	// 有调度者来进行日志写入，每条日志携带文件名
	lg = zap.New(core, zap.AddCaller())

	// 以上写入完成后，需要重新初始化，以便下一次日志的写入
	zap.ReplaceGlobals(lg)
	zap.L().Info("info logger success!")  // 记录写入日志成功
	return
}

// getLogWrite 传入logger配置，并获取日志写入器对象，用于写入日志
func getLogWrite(filename string, MaxAge, maxBackups, maxSize int) zapcore.WriteSyncer{
	// 实例化，获取日志对象
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxAge:     MaxAge,
		MaxBackups: maxBackups,
	}
	return zapcore.AddSync(lumberJackLogger)  // 返回日志写入器
}

// getEncoder 将配置文件中加载的配置，转换为固定的编码格式
func getEncoder() zapcore.Encoder{
	// 得到编码配置对象
	encoderConfig := zap.NewProductionEncoderConfig()
	// 设置ISO编码时间
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	// 设置级别编码
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 设置缓冲时间
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	// 设置调用者
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// 返回编码后的配置对象
	return zapcore.NewJSONEncoder(encoderConfig)
}

// GinLogger 接收gin框架默认的日志

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志记录
