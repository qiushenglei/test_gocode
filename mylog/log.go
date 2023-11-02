package mylog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
)

func getWriteSync() zapcore.WriteSyncer {

	// 文件输出
	file, _ := os.OpenFile("log.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	fileSync := zapcore.AddSync(file)

	// 标准输出(命令行)
	stdSync := zapcore.AddSync(os.Stdout)

	return zapcore.NewMultiWriteSyncer(fileSync, stdSync)
}

func getEncoder() zapcore.Encoder {
	encoderCfg := zap.NewProductionEncoderConfig()
	return zapcore.NewJSONEncoder(encoderCfg)
}

func Logger() {
	// 获取序列化格式
	enc := getEncoder()

	// 获取写入方式
	ws := getWriteSync()

	// newCore
	core := zapcore.NewCore(enc, ws, zap.DebugLevel)

	// 创建zap对象
	logger := zap.New(core, zap.AddCaller())
	sugarLogger := logger.Sugar()

	// 写
	_, trace, line, _ := runtime.Caller(0)
	sugarLogger.Error("add log", line, trace)

}
