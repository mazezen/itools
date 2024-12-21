package itools

import (
	"github.com/natefinch/lumberjack"
	"github.com/robfig/cron"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func NewLogger(fileName string, maxSize int, maxAge int, localTime, compress bool) *zap.Logger {
	encoder := zapEncoder()
	writeSyncer := zapWriteSyncer(fileName, maxSize, maxAge, localTime, compress)
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	consoleDebug := zapcore.Lock(os.Stdout)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	p := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.DebugLevel
	})
	var Codes []zapcore.Core
	Codes = append(Codes, core)
	Codes = append(Codes, zapcore.NewCore(consoleEncoder, consoleDebug, p))
	c := zapcore.NewTee(Codes...)
	return zap.New(c, zap.AddCaller())
}

func zapEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func zapWriteSyncer(fileName string, maxSize int, maxAge int, localTime, compress bool) zapcore.WriteSyncer {
	logger := &lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   maxSize,
		MaxAge:    maxAge,
		LocalTime: localTime,
		Compress:  compress,
	}
	c := cron.New()
	_ = c.AddFunc("0 0 0 1/1 * ?", func() {
		_ = logger.Rotate()
	})

	c.Start()
	return zapcore.AddSync(logger)
}
