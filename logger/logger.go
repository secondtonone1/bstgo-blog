package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Sugar *zap.SugaredLogger = nil

func getLogWriter() zapcore.WriteSyncer {
	// 日志文件
	//file, _ := os.Create("./log/blog.log")
	// file, err := os.OpenFile("./log/blog.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	// if err != nil {
	// 	panic("os.Openfile failed")
	// }
	// // 写日志
	// sync := zapcore.AddSync(file)

	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./log/blog.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}

	return zapcore.AddSync(lumberJackLogger)
}

func init() {
	// 编码器配置
	config := zap.NewProductionEncoderConfig()
	// 指定时间编码器
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	// 日志级别用大写
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	// 编码器
	encoder := zapcore.NewConsoleEncoder(config)
	writeSyncer := getLogWriter()
	// 创建Logger
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	Sugar = logger.Sugar()
	// 打印日志
	Sugar.Info("logger init success")

	// // SugaredLogger
	// sLog := logger.Sugar()
	// // 类似fmt.Println
	// sLog.Info("hello ", "yw, ", "helloZap")
	// // 类似fmt.Printf
	// sLog.Infof("hello %v, helloZap", "yw")

}

// func Println(arg ...interface{}) {
// 	sugar.Info(arg)
// }

// func Printf(template string, arg ...interface{}) {
// 	sugar.Infof(template, arg)
// }
