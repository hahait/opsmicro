package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	log *zap.Logger
)

func LogInit() {
	zenconfig := zap.NewProductionEncoderConfig()
	zenconfig.TimeKey = "time"
	zenconfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zenconfig.EncodeLevel = zapcore.CapitalLevelEncoder
	zencoder := zapcore.NewJSONEncoder(zenconfig)
	lj := &lumberjack.Logger{
		Filename:   "logs/user.log",
		MaxSize:    10240,
		MaxAge:     7,
		MaxBackups: 10,
		LocalTime:  true,
		Compress:   true,
	}
	zwsync := zapcore.AddSync(lj)
	zcore := zapcore.NewCore(zencoder, zwsync, zapcore.InfoLevel)
	log = zap.New(zcore, zap.AddCaller(), zap.AddCallerSkip(1))
	defer log.Sync()
}

func GetLogger() *zap.Logger {
	return log
}