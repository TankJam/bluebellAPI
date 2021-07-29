package logger

import (
	"bluebellAPI/settings"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
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
func GinLogger() gin.HandlerFunc {
	// 请求过来，通过logger记录http请求数据
	return func(c *gin.Context) {
		startTime := time.Now()
		path := c.Request.URL.Path  // URL
		query := c.Request.URL.RawQuery  // query values
		c.Next()  // 进入下一个中间件
		cost := time.Since(startTime) // 获取消耗时间
		lg.Info(path,
			zap.Int("status", c.Writer.Status()), // 日志中加入 status
			zap.String("method", c.Request.Method), // 日志中加入 method
			zap.String("path", path), // 日志中加入 path
			zap.String("query", query), // 日志中加入 query
			zap.String("ip", c.ClientIP()), // 日志中加入 ip
			zap.String("user-agent", c.Request.UserAgent()), // 日志中加入 UserAgent
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()), // 日志中加入 errors
			zap.Duration("status", cost), // 日志中加入 消耗时间
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志记录
func GinRecovery(stack bool) gin.HandlerFunc{
	return func(c *gin.Context) {
		defer func() {
			// recover 捕获异常, 需要用panic去检查断开的链接
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				// 判断异常是否为链接断开
				if ne, ok := err.(*net.OpError); ok {
					// 判断是否为系统调度异常
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer"){
							brokenPipe = true
						}
					}
				}

				// 可以输出请求体信息，以便排错
				httpRequest, _ := httputil.DumpRequest(c.Request, false)

				if brokenPipe {
					// true代表是链接错误，打印日志
					lg.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					lg.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					lg.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}

				c.AbortWithStatus(http.StatusInternalServerError)  // 返回服务错误状态码 500
			}
		}()
		c.Next()  // 进入下一个中间件
	}
}